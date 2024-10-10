package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := new(sync.WaitGroup) // similar to doing // wg := &sync.WaitGroup{}
	// each func call will spin up one go routine if go statement was added
	wg.Add(1) // wg.Add adds to a counter // keep track of number goroutines we are spinning
	go hello("Dev", wg)

	time.Sleep(5 * time.Second)

	fmt.Println("some more imp stuff going on")
	wg.Wait() // this line would block until counter is not reset to 0
	fmt.Println("End of the main")

	//time.Sleep(2 * time.Second) // bad idea
}

func hello(name string, wg *sync.WaitGroup) {

	defer wg.Done() // decrement the counter by one when it is called
	time.Sleep(5 * time.Second)
	fmt.Println("Hello, " + name)
}
