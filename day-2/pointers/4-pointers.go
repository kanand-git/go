package main

import "fmt"

//The second implication of copying a pointer is that
//if you want the value assigned to a pointer parameter to still be there
//when you exit the function, you must dereference the pointer and set the value.
//If you change the pointer, you have changed the copy, not the original.
//Dereferencing puts the new value in the memory location pointed to by both
//the original and the copy.

func main() {
	x := 10

	updateVal(&x)
	fmt.Println(x)
}
func updateVal(px *int) {
	abc := 20
	px = &abc // we have updated the reference here of *px , so it no longer updates x value from the main function
	*px = 100 // updating abc not x from the main
}
