package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("end of the main")
}

func greet() {
	data := os.Args[1:]
	//name, age, marks
	if len(data) != 3 {
		log.Println("please provide, name, age, marks")
		//os.Exit(1) // quit the program
		return // stops the exec of the current func
	}
	name := data[0]
	ageString := data[1]
	marksString := data[2]

	//var err error // errors are simply string messages
	// default value of error is nil, which means no error

	age, err := strconv.Atoi(ageString)
	// err happened if err value is not nil, it has some kind of msg
	// always handle the error in the next line after the func call
	if err != nil {
		log.Println(err)
		return
	}

	marks, err := strconv.Atoi(marksString)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(name, age, marks)

}
