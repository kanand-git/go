package main

import (
	"fmt"
	"sync"
	"time"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv

func main() {
	wg := &sync.WaitGroup{}

	ch := make(chan int) // unbuffered chan // it has size zero
	//i := go dowork() // we can't accept the returned values from the goroutine
	wg.Add(1)
	go doWork(ch, wg)

	//recv over the channel
	fmt.Println(<-ch) // it would block until doWork doesn't send value to the channel
	time.Sleep(5 * time.Second)
	fmt.Println(<-ch)
	wg.Wait()
}

func doWork(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// sender would block until there is no receiver
	ch <- 20 // sending the value to the channel
	ch <- 30
	fmt.Println("done") // done would execute after 5 seconds as receiver is delaying its startup
}
