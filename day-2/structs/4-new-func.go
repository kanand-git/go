package main

import (
	"github.com/username/reponame/db"
)

func main() {
	c := db.NewConf("mysql", "3306", "some Info")
	c.Insert()
	c.Info = "new info"

}
