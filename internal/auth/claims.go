package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var vErr = errors.New("Invalid Token")

type jwtClaims struct {
	jwt.StandardClaims
}

func (c jwtClaims) Valid() error {
	now := time.Now().Unix()

	if ok := c.VerifyAudience(auth0Audience, true); !ok {
		return vErr
	}

	if ok := c.VerifyIssuer(auth0Domain, true); !ok {
		return vErr
	}

	if ok := c.VerifyExpiresAt(now, true); !ok {
		return vErr
	}

	if ok := c.VerifyIssuedAt(now, true); !ok {
		return vErr
	}

	if ok := c.VerifyNotBefore(now, true); !ok {
		return vErr
	}

	return nil
}
