package main

import (
	"context"
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"strings"
)

type claims struct {
	jwt.RegisteredClaims
	Roles []string `json:"roles"`
}

//export TOKEN=TOKEN_VALUE
//curl -H "Authorization: Bearer $TOKEN" http://localhost:8082/check

func main() {
	PublicPEM, err := os.ReadFile("pubkey.pem")
	if err != nil {
		// If there's an error reading the file, print an error message and stop execution
		log.Fatalln("not able to read pem file")
	}

	// Parse the read public key to RSA public key format
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(PublicPEM)
	if err != nil {
		// If there's an error parsing the public key, log the error and stop execution
		log.Fatalln(err)
	}
	a := &auth{publicKey: publicKey}
	m := mid{a: a}
	http.HandleFunc("/check", m.AuthMid(check))
	panic(http.ListenAndServe(":8082", nil))
}

// check is an endpoint handler function for HTTP requests.
// It validates user's authentication from the request context.
func check(w http.ResponseWriter, r *http.Request) {
	// Extract the context from the request
	ctx := r.Context()
	// Extract the value for the authKey. The returned value is type asserted to be jwt.RegisteredClaims
	// ok will be false if the type assertion is not successful that implies the user is not authenticated
	v, ok := ctx.Value(authKey).(claims)
	if !ok {
		// Return HTTP Status Unauthorized if not authenticated
		http.Error(w, "not authenticated", http.StatusUnauthorized)
		return
	}
	// If authenticated, print the subject of the jwt token
	fmt.Println("current user", v.Subject)
	// Send back a success message to the client
	fmt.Fprintln(w, "welcome to check")
}

// The mid struct represents a middleware that has an auth structure attached to it
type mid struct {
	a *auth
}

// Define a new type 'key' that is an int
type key int

// Define a value for 'authKey' of type 'key'
var authKey key = 123

// AuthMid is a method on the mid struct
// It returns an endpoint handler function which checks for the validity of a Bearer JWT token in the Authorization header
func (m *mid) AuthMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Takes the "Authorization" header from the request
		authHeader := r.Header.Get("Authorization")
		// Splits the header into two parts separated by space
		// The first part should be "Bearer" and the second part should be the JWT token
		parts := strings.Split(authHeader, " ")

		// If the header doesn't contain exactly two parts or the first part is not "Bearer",
		// it will return an unauthorized status
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			http.Error(w, "expected authorization header format: Bearer <token>", http.StatusUnauthorized)
			return
		}

		// Try to validate the JWT token
		c, err := m.a.ValidateToken(parts[1])
		if err != nil {
			// If the token validation fails, return an unauthorized status
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		// Get the request context
		ctx := r.Context()
		// Add the JWT RegisteredClaims to the context with the authKey
		ctx = context.WithValue(ctx, authKey, c)

		// Call the next function in the chain with the enriched context
		next(w, r.WithContext(ctx))
	}
}

// The 'auth' struct has a publicKey of *rsa.PublicKey type which can be used to validate a JWT token
type auth struct {
	publicKey *rsa.PublicKey
}

// ValidateToken is a method on the 'auth' struct.
// It takes a JWT token string as input and returns the jwt.RegisteredClaims if the token is valid and an error otherwise.
func (a *auth) ValidateToken(token string) (claims, error) {
	var c claims
	// Parse the JWT token string and validate it with the publicKey stored in the 'auth' struct
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})

	// Return error if parsing the JWT token fails
	if err != nil {
		return claims{}, err
	}
	// Return error if the JWT token is not valid
	if !tkn.Valid {
		return claims{}, err
	}

	// Return the RegisteredClaims and nil error if the JWT token is successfully validated
	return c, nil
}
