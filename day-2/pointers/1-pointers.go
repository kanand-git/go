package main

import "fmt"

func main() {
	a := 10
	fmt.Println(&a) // & give access to the memory address of a variable
	p := &a         // copy the address of a to p
	*p = 20         // * is a dereference operator, access the value at the address
	fmt.Println(a)
	fmt.Println(p)

}
