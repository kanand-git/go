package main

import (
	"net/http"
	"small-app/handlers"
	"small-app/models"
)

func setupRoutes() {
	s := models.NewService("postgres")

	//setting the controller struct with
	c, err := handlers.NewController(s)
	if err != nil {
		panic(err)
	}
	// /user?user_id=123 // query
	http.HandleFunc("/user", c.GetUser)
	panic(http.ListenAndServe(":8080", nil))
}
