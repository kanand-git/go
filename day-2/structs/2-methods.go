package main

import "fmt"

// when to use pointer receiver or not
// https://go.dev/doc/faq#methods_on_values_or_pointers

type Student struct {
	Name string
	Age  int
}

func (s *Student) SayHello() { // func (receiver) MethodName(params) returnTypes
	//receiver is the var of the struct
	fmt.Println("Hello, my name is", s.Name)
}
func (s *Student) UpdateAge(age int) {
	s.Age = age
}

func main() {
	s := Student{"John", 20}
	s.SayHello()
	s.UpdateAge(21) // we are passing the address of s to the updateAge method
	fmt.Println(s.Age)
}
