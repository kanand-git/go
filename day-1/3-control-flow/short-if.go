package main

import "fmt"

func main() {
	// in short if we can declare variable in the if block and do comparisons on them
	// we can reuse the variable as a different type in another if block, because
	// the variable doesn't exist after the if block is over
	if a, _ := fmt.Println(); a > 10 { // _ could be used to ignore return values from the func

	}

	if a := fmt.Sprintf("hello %s", "dev"); a == "dev" {

	} else {

	}

}
