package main

import "fmt"

func main() {
	//var i interface{}
	var i any
	i = 10

	i = "hello"

	x, ok := i.(int) // type assertion // checking if interface is storing integer
	if !ok {
		fmt.Println("not an integer")
	} else {
		fmt.Println(x)
	}

	i = true
	i = 10.5
	i = []int{1, 2, 3}
	i = map[string]int{"a": 1, "b": 2}
	i = struct {
		a int
		b string
	}{
		a: 1,
		b: "hello",
	}
	fmt.Println(i)

}
