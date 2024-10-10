package main

import (
	"fmt"
	"sync"
)

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
var wg = &sync.WaitGroup{}

func main() {
	c := make(chan int)
	wg.Add(4)
	go addNum(10, 20, c)
	go mult(10, 10, c)
	go sub(100, 90, c)
	go calcAll(c)

	wg.Wait()

}

func addNum(a, b int, c chan int) {
	defer wg.Done()
	sum := a + b
	// send this to channel
	// in case of an unbuffered chan , receiver must be ready otherwise send will block
	// send operation signal on the channel  // signaling with data
	c <- sum
}

func sub(a, b int, c chan int) {
	defer wg.Done()
	diff := a - b

	// send
	c <- diff
}

func mult(a, b int, c chan int) {
	defer wg.Done()
	prod := a * b

	//send
	c <- prod
}

func calcAll(c chan int) {
	defer wg.Done()
	//TODO:// recv all the values from channel c
	x, y, z := <-c, <-c, <-c
	fmt.Println(x, y, z)
}
