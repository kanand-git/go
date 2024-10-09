package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

// Err should be prefixed to error msgs created using errors package or fmt package

var ErrFileNotFound = errors.New("in root directory")

func main() {
	// Attempt to open a file named "somefile.txt"
	f, err := OpenFile("somefile.txt")
	//err value would be wrapped errors from open file
	//[file not found,in root directory]
	if err != nil {

		// using errors.Is// we are looking inside the chain
		//[file not found,in root directory] // chain of errors
		//if it is present in the chain errors.Is would return true
		if errors.Is(err, ErrFileNotFound) { // Check if the error is a file not found error
			log.Println(err) // logging actual error to know what happend

			// we will take some compensating actions to fix the issue // not required everytime
			fmt.Println("Trying to create a file")

			// If the file is not found, attempt to create it
			f, err := os.Create("somefile.txt")
			// if it still fails we will quit it
			if err != nil {
				fmt.Println("File creation operation failed")
				return
			}
			defer f.Close() // Ensure the file is closed when the function returns
			return
		}
		// If another error occurred, log it and return // or error is not present in the chain
		log.Println(err)
		return
	}
	defer f.Close() // Ensure the file is closed when the function returns
}

// OpenFile attempts to open a file and returns a custom error if the file does not exist
func OpenFile(fileName string) (*os.File, error) {
	// Attempt to open the file
	f, err := os.Open(fileName)
	if err != nil {
		// Check if the error indicates that the file does not exist
		if errors.Is(err, os.ErrNotExist) {
			// Return a custom error wrapping the original error
			// wrapping an error put inside it an error chain that could be later inspected using errors.Is()
			return nil, fmt.Errorf("%w %w", err, ErrFileNotFound)
		}
		/*
				os.Open -> file not found
				ErrFileNotFound -> in root directory
			with %w we are wrapping errors together
			[file not found,in root directory]
		*/
		// Return any other error encountered
		return nil, err
	}
	// Return the opened file
	return f, nil
}
