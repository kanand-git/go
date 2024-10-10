package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)

	// NOTE:= if using buffered channel, don't use for select, this pattern would simply fail
	// use seperate go routines to receive the values, and close the channel after sending ,
	//in case of buffered channel
	c1 := make(chan int)
	c2 := make(chan int)
	c3 := make(chan int)
	done := make(chan struct{})

	wgWorker.Add(3)
	go func() {
		defer wgWorker.Done()
		//time.Sleep(2 * time.Second)
		c1 <- 1
	}()

	go func() {
		defer wgWorker.Done()
		//time.Sleep(1 * time.Second)
		c2 <- 2
	}()
	go func() {
		defer wgWorker.Done()
		c3 <- 3
		c3 <- 4

	}()

	go func() {
		wgWorker.Wait()
		close(done)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			// whichever case is not blocking exec that first
			//whichever case is ready first, exec that.
			// possible cases are chan recv , send , default
			case x := <-c1:
				fmt.Println(x)
				time.Sleep(1 * time.Second)
			case x := <-c2:
				fmt.Println(x)
				time.Sleep(1 * time.Second)
			case x := <-c3:
				fmt.Println(x)
				time.Sleep(1 * time.Second)
			case <-done:
				return
			}
		}
	}()
	//fmt.Println(<-c1)
	//fmt.Println(<-c2)
	//fmt.Println(<-c3)
	wg.Wait()
}
