package main

import "fmt"

// https://go.dev/ref/spec#Appending_and_copying_slices
/*
	append func working

	i := []int{10, 20, 30, 40, 50 } // len = 5 , cap =5
	append(i,60) // not enough cap so allocation is going to happen

//  sufficiently large underlying array.
	underlying array -> [10 20 30 40 50,60,{},{}] len =6 cap = 8

*/
func main() {
	// append helps to grow the slice
	var a []int = []int{10, 20, 30}
	//len : slice is referring to in the backing array,
	//cap : number of elms the slice can store
	fmt.Println("len = ", len(a), "cap= ", cap(a))
	fmt.Println(&a[0])

	//this append allocates new memory as cap was not sufficient
	a = append(a, 40)
	fmt.Println(&a[0])

	// this append doesn't create a new mem, enough cap present
	a = append(a, 50)
	fmt.Println(&a[0])

	fmt.Println("len = ", len(a), "cap= ", cap(a))
	fmt.Println(a)
}
