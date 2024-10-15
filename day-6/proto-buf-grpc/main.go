package main

import (
	"fmt"
	"proto-buf-grpc/proto"
)

func main() {
	simpleMessage()
}
func simpleMessage() {

	r := proto.BlogRequest{
		BlogId:  101,
		Title:   "Introduction to Protocol Buffers",
		Content: "Test",
	}

	fmt.Println(r.GetBlogId(), r.GetContent())

	fmt.Println(r.String())
}
