package main

import (
	"fmt"
)

// creating the new custom type named operation
type operation func(int, int) int

func main() {
	operate(sub, 10, 20)
	//http.HandlerFunc()
}

func sub(x, y int) int {
	return x - y
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op operation, x, y int) {
	sum := op(x, y)
	fmt.Println(sum)
}
