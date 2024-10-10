package main

import (
	"fmt"
	"sync"
)

// A send on a buffered channel can proceed if there is room in the buffer.
func main() {
	wg := new(sync.WaitGroup)

	// Define a buffered channel that can hold 2 integers
	ch := make(chan int, 2)
	wg.Add(1)
	go func() {
		defer wg.Done()
		ch <- 20
		ch <- 30
		ch <- 40 // it blocks until recv takes out the value
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch)
	}()

	wg.Wait()

}
