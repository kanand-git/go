package main

import (
	"fmt"
	"log"
)

func main() {
	// _ is used to ignore return values from the function
	//name, _ := hello("ajay", 33, 88)
	msg, ok := hello("ajay", 33, 88)
	if !ok { // ok == false
		log.Println("process failed", msg)
		return
	}
	fmt.Println("process success", msg)
}

// hello function return some msg, and return some bool value
func hello(name string, age, marks int) (string, bool) {
	if name == "" {
		return "please provide a name", false // it will stop the current func and return values
	}
	if age == 0 {
		return "please provide your age", false
	}
	if marks == 0 {
		return "please provide your marks", false
	}
	fmt.Println(name, age, marks)
	return "success", true
}
