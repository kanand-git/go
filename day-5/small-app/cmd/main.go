package main

import (
	"log"
	"net/http"
	"small-app/handlers"
	"small-app/models"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// it would initialize the conn once at the start of the app
	s := models.NewService("postgres")
	c, err := handlers.NewController(s)
	if err != nil {
		panic(err)
	}

	// /user?user_id=123 // query
	http.HandleFunc("/user", c.GetUser)
	panic(http.ListenAndServe(":8080", nil))
}
