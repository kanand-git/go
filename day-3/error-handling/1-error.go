package main

import (
	"errors"
	"fmt"
	"log"
)

var user = make(map[int]string)

// Err should be prefixed to error msgs created using errors package or fmt package

var ErrNotFound = errors.New("not found")

//var ErrNotFound = fmt.Errorf("not found %v", "some value")

func main() {
	//os.ErrClosed
	s, err := FetchRecord(1)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(s)
}

func FetchRecord(id int) (string, error) {
	name, ok := user[id]
	if !ok {
		return "", ErrNotFound
	}
	return name, nil
}
