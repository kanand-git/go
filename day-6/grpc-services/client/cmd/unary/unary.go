package main

import (
	pb "client/gen/proto" // renaming the import to pb
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
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

	// setting up required things to do the request
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	// Create a SignupRequest object
	req := &pb.SignupRequest{
		User: &pb.User{
			Name:     "John",
			Email:    "john@email.com",
			Password: "abc",
			Roles:    []string{"ADMIN", "USER"},
		},
	}
	resp, err := client.Signup(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Result)

}
