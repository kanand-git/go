package main

import (
	"fmt"
	"os"
)

func main() {
	// 0 index would be the location to the binary, followed by the args user would pass
	fmt.Println(os.Args[1:])

}
