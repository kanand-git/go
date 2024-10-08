package db

import "fmt"

type Conf struct {
	dbName string // unexported, no one can change it outside the current package
	Info   string // which we don't care to be changed
}

// NewConf returns an object or instance of the conf struct,
// in this we intialize unexported fields of the conf struct
func NewConf(dbName, portNumber string, info string) Conf {
	return Conf{
		dbName: dbName + ":" + portNumber,
		Info:   info,
	}
}
func (c Conf) Insert() {
	fmt.Println("inserting in the ", c.dbName)
}
