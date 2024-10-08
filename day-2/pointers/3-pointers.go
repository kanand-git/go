package main

import "fmt"

func main() {
	//a := 10
	//var x *int = &a
	//var b *int = &a
	//z := 20
	//b = &z

	var p *int     // default value of pointer is nil
	updateValue(p) // passing nil to the updateValue
	// value of p would be copied to updateValue func param

	// p value is nil, we cant dereference nil memory, hence panic
	fmt.Println(*p) // dereference the pointer and access the value
}

func updateValue(p1 *int) { // it would receive nil from the main
	x := 10
	p1 = &x          // storing the address of x in p1 // updating p1 from nil to let say x80(address)
	fmt.Println(*p1) // dereference the value and prints it.
}
