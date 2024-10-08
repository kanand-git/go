package main

import "fmt"

func main() {
	//
	operate(add, 10, 20)
	operate(sub, 10, 20)

}
func someStuff(abc func(s string) bool) {

}

// function signature is datatype of that function in go

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(x int, y int) int, s, y int) {
	sum := op(s, y)
	fmt.Println(sum)
}
