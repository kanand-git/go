package main

import (
	"fmt"
	"sync"
)

var wg = new(sync.WaitGroup)

func main() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i)
	}

	wg.Wait()
}

func work(id int) {
	defer wg.Done()
	// anonymous func, a function without name
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(id)
	}() // () this is calling of the function

}
