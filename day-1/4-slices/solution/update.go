package main

import "fmt"

func main() {
	x := []int{10, 20, 30, 40, 50, 60, 70}

	update(x)
	fmt.Println(x)
	x = appendValues(x)
	fmt.Println(x)
}

// if you just want to update, **not append**, then you can pass the slice directly to the func
// and do your changes, and that works
func update(s []int) {
	s[0] = 100
}

// if you want to append
// if you are appending the values to the slice, always return the slice
// returning the slice would ensure that caller has updated reference
func appendValues(s []int) []int {
	s = append(s, 100)
	return s
}
