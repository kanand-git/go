// package name should be same as the folder name
// lib package -> package that you can't run , we can still use go build for compile time errors

package sum

import (
	"fmt"
)

// in one folder only one package can exist, we cant create two different package declarations

// to export the func , make the first letter uppercase
// the func can be used outside of the current package, in our case sum

var Total int

func Addition(a, b int) {
	Total = a + b
	sub() // it is the part of the same package , so it can be called directly
	fmt.Println("i am trying to add values", Total)
}
