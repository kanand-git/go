package main

import "fmt"

func main() {
	defer fmt.Println(1) // the line deferred for exec
	// the function call would be done when surrounding func returns
	defer fmt.Println(2)
	// defer maintains a stack, last in first out

	fmt.Println(3)
	panic("some serious problem")
	fmt.Println(4)
}
