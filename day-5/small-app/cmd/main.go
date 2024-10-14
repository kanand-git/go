package main

import (
	"log"
)

// main-> if request hits -> handlers -> models
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	// it would initialize the conn once at the start of the app
	setupRoutes()
}
