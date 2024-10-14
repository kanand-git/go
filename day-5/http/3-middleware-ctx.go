package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// go get moduleName / go get "github.com/google/uuid"
// you must have a module setup in the current project for this to work

// Creating a custom type for context key
type reqKey int

// A constant for request id key
const RequestIDKey reqKey = 123

func main() {
	http.HandleFunc("/home", RequestIdMid(LoggingMid(homeP)))
	panic(http.ListenAndServe(":8080", nil))
}

// Function to handle requests at homePage
func homeP(w http.ResponseWriter, r *http.Request) {
	// Print logs for each request received
	ctx := r.Context()
	// type assertion // checking if the value is still of the string type
	reqId, ok := ctx.Value(RequestIDKey).(string)
	if !ok {
		reqId = "unknown"
	}

	log.Println(reqId, "In home Page handler")
	//ctx, cancel := context.WithTimeout(ctx, time.Second*10)

	FindUser(ctx, reqId)
	// Respond to the client request
	fmt.Fprintln(w, "this is my home page")
}

func RequestIdMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()
		// we should not create an empty container
		// ctx := context.Background()
		ctx := r.Context() // we will use the context already available in the request object
		// below line have updated context with requestId in it
		ctx = context.WithValue(ctx, RequestIDKey, id)

		// updating the request object with the updated ctx
		r = r.WithContext(ctx)

		next(w, r)
	}
}

func LoggingMid(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// type assertion // checking if the value is still of the string type
		reqId, ok := ctx.Value(RequestIDKey).(string)
		if !ok {
			reqId = "unknown"
		}
		// Log the details of the request
		log.Printf("%s : started   : %s %s ",
			reqId,
			r.Method, r.URL.Path)
		defer log.Printf("%s : completed : ", reqId)
		next(w, r)

	}
}

// ctx should be the first param in the function // this is a go convention

func FindUser(ctx context.Context, id string) {
	log.Println(id, "Finding user")
	db := &sql.DB{}
	r, err := db.ExecContext(ctx, "")
	_, _ = r, err
}
