package main

import (
	"fmt"
	"log"
	"strconv"
)

func SumString(s, x string) (int, error) { // err must be the last value to be returned
	a, err := strconv.Atoi(s)

	if err != nil {
		// if returning the err , then avoid logging to avoid duplicate logs
		//log.Println("err msg", err)
		return 0, err // whenever err happens set other values to default
	}

	b, err := strconv.Atoi(x)
	if err != nil {
		//log.Println("err msg", err)
		return 0, err
	}

	// success case
	return a + b, nil // don't write err , even err is going to be nil
}

func main() {
	sum, err := SumString("abc", "0")
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(sum)

}
