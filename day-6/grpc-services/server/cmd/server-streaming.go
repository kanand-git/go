package main

import (
	"log"
	pb "server/gen/proto"
	"time"
)

// Server Streaming
// CLient will send one single request, server will send multiple stream of responses back

func (u userService) GetPosts(req *pb.GetPostsRequest, stream pb.UserService_GetPostsServer) error {
	id := req.GetUserId()
	log.Println("GetPosts: ", "fetching all posts for user id ", id)
	//assume these posts we are getting in batches
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

	// preparing the first response
	res := &pb.GetPostsResponse{Posts: batch1}

	// sending the resp to the client
	err := stream.Send(res)
	if err != nil {
		log.Println("Error sending response to client: ", err)
		return err
	}
	time.Sleep(3 * time.Second)
	//constructing the second batch
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

	b2 := &pb.GetPostsResponse{Posts: batch2}

	err = stream.Send(b2)
	if err != nil {
		log.Println("Error sending response to client: ", err)
		return err
	}

	log.Println("all posts are sent for user id", id)
	// returning from function would mark the end of the stream and close the connection
	return nil

}
