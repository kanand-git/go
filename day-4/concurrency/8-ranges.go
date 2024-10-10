package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	// this waitgroup would track the number of worker goroutines spawned
	wgWorker := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			// running n number of task , all goroutines are pushing to the same channel
			// keeping track of number of worker goroutines spawned
			wgWorker.Add(1)
			go func() {
				defer wgWorker.Done()
				ch <- i
			}()
		}

		//we need to block our goroutine before closing the channel
		//because we want to make sure all the work
		// is done and finished // closing a channel will stop the for range loop
		wgWorker.Wait() // waiting until the worker goroutines are not finished
		close(ch)       //sending is finished over the channel ch
	}()

	for v := range ch {
		//ranging until the channel is not closed
		//range would receive all the remaining values even after the channel is closed
		fmt.Println(v)
	}

}
