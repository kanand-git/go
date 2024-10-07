package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	// creating a new backing array
	b := make([]int, len(x), cap(x))

	// src, dest
	copy(b, x)
	fmt.Println("b=", b)
	b[0] = 100 // this will not affect x, as b is using a different backing array to store the elems
	fmt.Println("x=", x)
	fmt.Println("b=", b)
}
