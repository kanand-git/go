package main

import (
	pb "client/gen/proto"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

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

	stream, err := client.GetPosts(context.Background(), &pb.GetPostsRequest{UserId: 1})
	if err != nil {
		log.Fatal(err)
	}

	// running infinite loop to recv all the responses from the stream
	for {
		post, err := stream.Recv()
		// using errors.Is to check if the stream has finished
		if errors.Is(err, io.EOF) {
			break
		}
		//any other error we will quit
		if err != nil {
			log.Println(err)
			return
		}
		select {
		// checking if server cancelled the request
		case <-stream.Context().Done():
			fmt.Println("server cancelled the request ")
			return
		default:
			// continue
		}
		fmt.Println(post)
		fmt.Println()

	}

}
