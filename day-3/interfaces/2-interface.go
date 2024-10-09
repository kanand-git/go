package main

import (
	"fmt"
	"log"
)

type user struct {
	name  string
	email string
}

func (u user) Write(p []byte) (n int, err error) {
	fmt.Printf("sending a notification to %s %s %s", string(p), u.name, u.email)
	return len(p), nil
}
func main() {
	u := user{name: "john", email: "john@gmail.com"}
	//l := log.New(os.Stdout, "sales: ", log.Lshortfile)
	// to pass user object log.New func we need to implement the interface because log.New accepts io.Writer type values
	l := log.New(u, "sales: ", log.Lshortfile)
	l.Println("hello")
}

//func New(out io.Writer, prefix string, flag int) *Logger

//New creates a new Logger. The out variable sets the destination to which log data will be written.
//The prefix appears at the beginning of each generated log line, or after the log header if the Lmsgprefix flag is provided.
//The flag argument defines the logging properties.

// type Writer interface {
//	Write(p []byte) (n int, err error)
//}
