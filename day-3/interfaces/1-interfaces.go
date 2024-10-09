package main

import (
	"fmt"
)

// Polymorphism means that a piece of code changes its behavior depending on the
// concrete data it’s operating on // Tom Kurtz, Basic inventor

// "Don’t design with interfaces, discover them". - Rob Pike

// if your function needs to support multiple implementation at the start of the project, then above convention doesn't apply
// or if you want to write mocks for testing then also creating interfaces would be required

// interfaces are automatically implemented when a type implements all the methods of the interface

// Bigger the interface weaker the abstraction // Rob Pike

type Reader interface {
	Read(b []byte) (int, error)
}

type file struct {
	name string
}

func (f file) NotAnInterfaceMethod() {
	fmt.Println("not an interface method of files")
}

func (f file) Read(b []byte) (int, error) {
	fmt.Println("reading files and processing them", f.name)
	return 0, nil
}

type json struct {
	data string
}

func (j json) Read(b []byte) (int, error) {
	fmt.Println("reading json", j.data)
	return 0, nil
}

// DoWork can accept any type that implements the interface
func DoWork(r Reader) {
	b := make([]byte, 1024)
	r.Read(b)

	// type assertion // checking if file struct is present in the interface and doing some file specific task
	x, ok := r.(file)
	if ok {
		x.NotAnInterfaceMethod()
	}
}

func main() {
	f := file{"test.txt"}
	// `` used for raw strings, no processing inside string
	j := json{`{"name": "test"}`}
	DoWork(f)
	DoWork(j)

}
