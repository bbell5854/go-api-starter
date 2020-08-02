package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	jwtware "github.com/gofiber/jwt"
)

const auth0Audience = "https://scrumdelta.auth0.com/api/v2/"
const auth0Domain = "https://scrumdelta.auth0.com/"

type jwks struct {
	Keys []jsonWebKeys `json:"keys"`
}

type jsonWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func Protected() func(c *fiber.Ctx) {
	signingKeys, err := getSigningKeys()
	if err != nil {
		log.Printf("Couldn't fetch signing keys: %s", err)
	}

	Claims := jwtClaims{}

	t := reflect.ValueOf(Claims).Type().Elem()
	// claims := reflect.New(t).Interface().(jwt.Claims)
	json, _ := json.Marshal(t)
	fmt.Println(string(json))
	// token, err = jwt.ParseWithClaims(auth, claims, cfg.keyFunc)

	return jwtware.New(jwtware.Config{
		SigningMethod: jwt.SigningMethodRS256.Name,
		SigningKeys:   signingKeys,
		ErrorHandler:  errorHandler,
	})
}

func getSigningKeys() (map[string]interface{}, error) {
	signingKeys := make(map[string]interface{})

	resp, err := http.Get(fmt.Sprintf("%s.well-known/jwks.json", auth0Domain))
	if err != nil {
		return signingKeys, err
	}
	defer resp.Body.Close()

	var jwks = jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)
	if err != nil {
		return signingKeys, err
	}

	for i, jwksKey := range jwks.Keys {
		strI := strconv.Itoa(i)
		signingKeys[strI] = jwksKey.Kid
	}

	return signingKeys, nil
}

func errorHandler(c *fiber.Ctx, err error) {
	fmt.Println("HOTDOG")
	c.SendStatus(http.StatusUnauthorized)
}
