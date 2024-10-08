package main

import (
	"fmt"
)

func main() {
	num := 50 // &num = x80
	p := &num // ptr address = x100, ptr value = x80
	fmt.Println("Before modification:", *p)
	modifyDoublePointers(&p)

	// The value was changed, and the pointer re-assignment succeeded
	fmt.Println("After modification:")
	fmt.Println("Value pointed by ptr:", *p)

	fmt.Println("Pointer address:", p)
	fmt.Println("Value pointed by ptr:", num)
	//json.Unmarshal()
	//errors.As()

}

// double pointer stores the pointer address,
func modifyDoublePointers(ptr **int) {
	// Modifying the actual integer value
	**ptr = 100
	// Reassigning the pointer itself
	newValue := 200
	*ptr = &newValue
}
