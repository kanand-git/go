package main

import (
	"database/sql"
	"fmt"
)

func main() {

	doWork()
	fmt.Println("mission critical stuff, it doesn't depends on the updateList")
}

func doWork() {
	// RecoverPanic would recover the current function from panic, but the function needs to stop
	// it can't continue executing
	defer RecoverPanic()
	updateList([]int{})
	fmt.Println("work is done")
}
func updateList(s []int) {
	// defer guarantees to run // so it would recover the panic if it would happen
	s[0] = 100
}

func RecoverPanic() {
	// The built-in `recover` function can stop the process of panicking,
	//if it is called within a deferred function.

	// msg would have the actual panic message if that happened
	msg := recover()

	if msg != nil {
		// If `recover` captured a panic, it returns the panic value.
		// Here we print it.
		fmt.Println(msg)
		//fmt.Printf("%s", debug.Stack())
	}
}

func WeCanHavePanic() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		//Don't Panic until it is crucial to run your app
		panic(err)
	}
	defer db.Close()
}
