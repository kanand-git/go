package main

import "fmt"

func main() {

	//var i [5]int = [5]int{10, 20, 30} // array // arrays are fixed size
	i := [5]int{10, 20, 30}
	i[0] = 100

	// slices are pointers
	// slices point to array in the memory which is known as backing/underlying array
	//var x []int // default value of slice is nil

	//make creates a backing array and slice points to it
	x := make([]int, 0, 2) // type, len, cap
	fmt.Println(i)
	fmt.Printf("%#v\n", x)

	if x == nil {
		fmt.Println("it is nil slice")
		return
	}

	// panic, because x[0] is not present
	x[0] = 100 // this is to update, not to add
}
