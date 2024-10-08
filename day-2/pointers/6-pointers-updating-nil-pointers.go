package main

import "fmt"

func main() {
	var p *int
	updateValue(&p)
	fmt.Println(*p) // p in no longer nil, as we updated its value to store new address
}

func updateValue(p1 **int) {
	x := 10
	*p1 = &x // updating the pointer from the main function , to store a new address
	//fmt.Println(*p1)
}
