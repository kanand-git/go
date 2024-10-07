package main

import (
	"fmt"
)

// go mod init moduleName // initialize a go module
// it helps in managing packages and external deps

//go run 1-var.go
//go build 1-var.go
//./1-var

// go is a statically compiled language

// global scope
var someData string

// a:= 10 // shorthand doesn't work in global scope
func main() {
	// in local scope all the variables must be used
	var s string // default value string ""
	var i = 10
	//i = "hello" // go is a statically compiled language
	fmt.Println("hello everyone")

	fmt.Println(s, i)

	sName, sAge := "raj", 12 // shorthand operator // it creates and assign the value

	fmt.Println(sName, sAge)

	var (
		//camelCase
		uName  string
		uAge   int = 15
		uMarks float64
	)

	fmt.Println(uName, uMarks, uAge)
	//time.Second // peek into it for design pattern
	//http.StatusInternalServerError
	//os.O_APPEND

}
