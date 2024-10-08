package main

import "fmt"

func main() {
	//this function would not be called at below line
	//we are defining the working of the func
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	// calling of the func
	add(1, 2)
	_ = add // quick hack to avoid errors for unused var, not to be used in prod code
}
