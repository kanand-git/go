package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"os"
	"time"
)

func main() {
	privatePem, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(privatePem)
	if err != nil {
		log.Fatal(err)
	}
	// iss (issuer): Issuer of the JWT
	// sub (subject): Subject of the JWT (the user)
	// aud (audience): Recipient for which the JWT is intended
	// exp (expiration time): Time after which the JWT expires
	// nbf (not before time): Time before which the JWT must not be accepted for processing
	// iat (issued at time): Time at which the JWT was issued; can be used to determine age of the JWT
	// jti (JWT ID): Unique identifier; can be used to prevent the JWT from being replayed (allows a token to be used only once)
	claims := struct {
		jwt.RegisteredClaims
		Roles []string
	}{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "api project",
			Subject:   "101",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(50 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Roles: []string{"USER"},
	}

	tkn := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	str, err := tkn.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
