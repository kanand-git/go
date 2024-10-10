package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func slowFunction() int {
	time.Sleep(5 * time.Second)
	fmt.Println("slow fn ran and add 100 records to db")
	fmt.Println("receiver should process it")
	return 42
}

func main() {
	ch := make(chan int, 2) // Creating a buffered channel of integers with a capacity of 2

	// Creating a context with a timeout of 2 seconds. It gives us an empty container to put timer values.
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel() // Ensures that the resources taken up by the time package are cleaned up

	wg := &sync.WaitGroup{} // Creating a wait group to sync the completion of goroutines

	// Adding a counter for the sender goroutine to the wait group
	wg.Add(1)
	// Launching the sender goroutine
	go func() {
		defer wg.Done() // Decrementing the counter in the wait group when the goroutine completes

		// A slow function call that takes some time to produce a result
		x := slowFunction()
		select {
		// Trying to send the result to the channel
		case ch <- x:
			fmt.Println("value sent to channel")
		// If the context is done (timeout or cancel), handle accordingly
		case <-ctx.Done():
			fmt.Println("receiver is not present, doing rollback logic")
		}
	}()

	// Adding a counter for the receiver goroutine to the wait group
	wg.Add(1)
	// Launching the receiver goroutine
	go func() {
		defer wg.Done() // Decrementing the counter in the wait group when the goroutine completes
		select {
		// If the context is done (timeout), print the error and return
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		// If a value is received from the channel, process it
		case x := <-ch:
			fmt.Println("value received", x)
			fmt.Println("all the records processed")
		}
	}()

	// Waiting for all goroutines in the wait group to finish
	wg.Wait()
}
