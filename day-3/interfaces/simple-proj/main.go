package main

import (
	"simple-proj/stores"
	"simple-proj/stores/mysql"
	"simple-proj/stores/postgres"
)

func main() {
	u := stores.User{
		Name:  "ajay",
		Email: "ajay@email.com",
	}
	// create connection to mysql
	m := mysql.New("mysql conn")
	// create connection to postgres
	p := postgres.New("postgres conn")

	// call stores.NewService
	ms := stores.NewService(m)
	ps := stores.NewService(p)

	// Call Create method of mysql and postgres using interface which is inside service struct
	ms.Create(u)
	ps.Create(u)

}
