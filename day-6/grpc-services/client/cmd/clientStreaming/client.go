package main

import (
	pb "client/gen/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

// CLientStreaming // Client would send multiple request to the server and server would return one single response
func main() {
	dialOpts := []grpc.DialOption{
		// WithTransportCredentials specifies the transport credentials for the connection
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	// creaing grpc client, it will connect to remote server on specified port
	conn, err := grpc.NewClient("localhost:5001", dialOpts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// creating instance of the client struct to call remote methods
	client := pb.NewUserServiceClient(conn)

	// Create a context with a 10-second timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Call the CreatePost method of the client and get a reply stream, along with any error
	stream, err := client.CreatePost(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	// Simulate first batch of posts
	batch1 := []*pb.Post{
		{
			Title:  "The Science of Design",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "The Politics of Power",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "The Art of Programming",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	req := &pb.CreatePostRequest{Posts: batch1}

	// Attempt to send CreatePost request through the stream
	err = stream.Send(req)
	if err != nil {
		log.Println(err)
		return
	}

	time.Sleep(5 * time.Second)
	batch2 := []*pb.Post{
		{
			Title:  "Post 11",
			Author: "Author 1",
			Body:   "Body of post 1",
		},
		{
			Title:  "Post 21",
			Author: "Author 2",
			Body:   "Body of post 2",
		},
		{
			Title:  "Post 31",
			Author: "Author 3",
			Body:   "Body of post 3",
		},
	}

	// Put the second batch in a CreatePostRequest object
	req = &pb.CreatePostRequest{Posts: batch2}

	// Attempt to send CreatePost request through the stream
	err = stream.Send(req)
	if err != nil {
		log.Println(err)
		return
	}

	// Close the client streaming and receive the server's response
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(resp.Result)

}
