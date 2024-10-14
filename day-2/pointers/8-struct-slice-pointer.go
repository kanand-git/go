package main

import "fmt"

// Person struct with a names slice
type Person struct {
	names []string
}

// Method to append a value to the names slice

func (p *Person) AddName(name string) {
	p.names = append(p.names, name) // this would directly update the names slice
	// we don't have to return it
	// but if we are getting a slice in func param, we should always return it in case of append
}

func main() {
	// Create a new Person instance
	p := &Person{}

	// Append names using the AddName method
	p.AddName("Alice")
	p.AddName("Bob")
	p.AddName("Charlie")

	// Print the names slice
	fmt.Println("Names:", p.names)
}
