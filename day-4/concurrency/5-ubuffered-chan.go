package main

import "fmt"

// https://go.dev/ref/spec#Send_statements
// A send on an unbuffered channel can proceed if a receiver is ready.
// send will block until there is no recv
func main() {
	go addNum(10, 20, c)
	go mult(10, 10, c)
	go sub(100, 90, c)
	go calcAll(c)

}

func addNum(a, b int, c chan int) {
	//defer wg.Done()
	sum := a + b
	// send this to channel
	// in case of an unbuffered chan , receiver must be ready otherwise send will block
	// send operation signal on the channel  // signaling with data
}

func sub(a, b int, c chan int) {
	//defer wg.Done()
	diff := a - b

	// send
	c <- diff
}

func mult(a, b int, c chan int) {
	//defer wg.Done()
	prod := a * b

	//send
}

func calcAll(c chan int) {
	//TODO:// recv all the values from channel c

	fmt.Println(x, y, z)
}
