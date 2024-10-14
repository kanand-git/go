package main

import (
	"fmt"
	"net/http"
)

func main() {
	//Mid1 is a middleware that exec some preprocessing logic or post processing logic
	// a middleware accepts a handler func and returns a handler func
	// we are calling the Mid1 so the return type of Mid1 should satisfy the http.HandlerFunc
	http.HandleFunc("/home", Mid1(Mid2(homePage)))

	panic(http.ListenAndServe(":8080", nil))
}

func Mid1(next http.HandlerFunc) http.HandlerFunc {
	// we need to return a func from this
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 1 started")
		next(w, r) // handler function would be called here that was passed to mid
		fmt.Println("middleware 1 was completed")
	}

}

func Mid2(next http.HandlerFunc) http.HandlerFunc {
	// we need to return a func from this
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("middleware 2 invoked")
		next(w, r) // hander function would be called here that was passed to mid
		fmt.Println("middleware 2 was completed")
	}

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homePage invoked")
	w.Write([]byte("Hello World"))
}
