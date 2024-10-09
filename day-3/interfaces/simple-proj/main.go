package main

import (
	"simple-proj/stores"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	s := stores.NewService()
	s.Create()
	// create connection to mysql
	// create connection to postgres
	// call stores.NewService
	// Call Create method of mysql and postgres using interface which is inside service struct

}
