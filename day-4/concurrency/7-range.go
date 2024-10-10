package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		// close signal range that no more values be sent and it can stop after receiving remaining values
		// close the channel when sending is finished
		//close(ch) // we cant send values to the closed channel
		// we can recv the reaming values from the closed channel
		//ch <- 11 // this would panic because chan is closed
	}()

	//close(ch) // there is a chance where goroutine never ran and we closed the channel already

	// recv over the channel // it would run infinitely until the channel is close
	for v := range ch {
		fmt.Println(v)
	}
	wg.Wait()
}
