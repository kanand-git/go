package main

import "fmt"

func main() {
	// in this case we call a func, the func return value would be passed as parameter
	//for operate func in this case
	operate(add(), 10, 20)
	operate(sub, 10, 20)
	var s string = "hello world"
	s := hello()
	fmt.Println(s)

}

// return type of the function matches the op parameter in operate function
// that's why we can call add() in main and it would work
func add() func(int, int) int {
	return func(x int, y int) int {
		return x + y
	}
}

func sub(x, y int) int {
	return x - y
}

// operate func can accept function in op parameter,
// the function signature we are passing should match to op parameter type
func operate(op func(x int, y int) int, x, y int) {
	sum := op(x, y)
	fmt.Println(sum)
}

func hello() string {
	return "hello world"
}
