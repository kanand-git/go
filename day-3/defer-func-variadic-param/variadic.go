package main

import "fmt"

type conf struct {
	port int
}

func main() {
	PrintNames(1, 33.5, "a", "b")
}

// PrintNames is a variadic function
// variadic param can accept any number of values of the certain types
// it should be the last param in the func
func PrintNames(id int, marks float32, names ...string) {
	fmt.Printf("%T\n", names)
	//names[0] // this should not be used with variadic type value without checking the len
	//len(names) // check for len before doing any operation using index
	for _, name := range names {
		fmt.Println(name)
	}
	fmt.Println(id, marks)
}
