package main

import "fmt"

func main() {
	x := 10
	ptr := &x
	// go is always pass by value,
	update(ptr) // value of ptr would be copied to update func param
	fmt.Println(x)
}

// whenever you want to modify a param, instead of returning the updated value, use a pointer
func update(p *int) {
	*p = 20
}
