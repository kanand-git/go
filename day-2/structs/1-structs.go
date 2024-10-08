package main

import "fmt"

type Person struct {
	Name  string //fields
	Age   int
	marks []int
}

func main() {
	var p Person
	fmt.Println(p)
	fmt.Printf("%+v\n", p) // print field value pairs
	fmt.Printf("%#v", p)

	p.Name = "Rajesh"
	p.Age = 20
	p.marks = []int{100, 90, 80}

	p1 := Person{Name: "Rajesh", Age: 20, marks: []int{100, 90, 80}}
	fmt.Println(p1)
}
