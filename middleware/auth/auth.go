package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
)

const auth0Audience = "https://scrumdelta.auth0.com/api/v2/"
const auth0Domain = "https://scrumdelta.auth0.com/"

const authHeaderKey = "Authorization"
const authScheme = "Bearer"

// Jwks contains a slice of JSONWebKeys
type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

// JSONWebKeys contains our cert key information
type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func validationKeyGetter(token *jwt.Token) (interface{}, error) {
	if ok := token.Claims.(jwt.MapClaims).VerifyAudience(auth0Audience, false); !ok {
		return token, errors.New("invalid audience")
	}

	if ok := token.Claims.(jwt.MapClaims).VerifyIssuer(auth0Domain, false); !ok {
		return token, errors.New("invalid issuer")
	}

	cert, err := getPemCert(token)
	if err != nil {
		return nil, err
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func getPemCert(token *jwt.Token) (string, error) {
	var cert string

	resp, err := http.Get(fmt.Sprintf("%s.well-known/jwks.json", auth0Domain))
	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return cert, err
	}

	for k := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err = errors.New("unable to find appropriate key")
	}

	return cert, err
}

func getJwtFromHeader(c *fiber.Ctx) (string, error) {
	auth := c.Get(authHeaderKey)
	l := len(authScheme)

	if len(auth) > l+1 && auth[:l] == authScheme {
		return auth[l+1:], nil
	}

	return "", errors.New("Missing or malformed JWT")
}

func checkToken(c *fiber.Ctx) bool {
	token, err := getJwtFromHeader(c)
	if err != nil {
		fmt.Printf("Error: %v", err)
		return false
	}

	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, validationKeyGetter)
	if err != nil {
		fmt.Printf("Error parsing token: %v", err)
		return false
	}

	c.Locals("userSlug", claims["sub"])

	return parsedToken.Valid
}

// Protected does check your JWT token and validates it
func Protected() func(*fiber.Ctx) {
	return func(c *fiber.Ctx) {
		if ok := checkToken(c); !ok {
			c.SendStatus(http.StatusUnauthorized)
			return
		}

		c.Next()
	}
}
