package main

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	pb "server/gen/proto"
)

// CLientStreaming // Client would send multiple request to the server and server would return one single response

func (u userService) CreatePost(stream pb.UserService_CreatePostServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				// if the stream is finished from the client side, we would quite the loop not the function
				fmt.Println("stream is finished")
				break
			}
			log.Println(err)
			return err
		}

		select {
		case <-stream.Context().Done():
			log.Println("client closed stream")
			return status.Error(codes.Canceled, "client closed stream")
		default:
			// client is still connected
		}

		log.Printf("Received Create Post Requests: %v\n", req.GetPosts())
		//send the post for further processing, pass the context to the next call
		// if the ctx is cancelled , the other layer would also know and can rollback changes made if required
		// func (ctx, req.GetPost)
	}

	return stream.SendAndClose(&pb.CreatePostResponse{Result: "all posts created"})
}
