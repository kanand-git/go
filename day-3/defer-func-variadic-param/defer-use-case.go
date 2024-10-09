package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.Open("test.txt")
	//defer f.Close() // nil.Close() // if there was an error while opening file, connection would be nil
	// first step after calling a func is always error handling
	if err != nil {
		// os.exit
		log.Fatal(err)
	}
	defer f.Close()

}
