package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"server/gen/proto"
)

type userService struct {
	proto.UnimplementedUserServiceServer
}

func main() {
	listner, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Println(err)
		return
	}

	//NewServer creates a gRPC server which has no service registered
	//and has not started to accept requests yet.
	s := grpc.NewServer()
	proto.RegisterUserServiceServer(s, &userService{})

	//exposing gRPC service to be tested by postman
	reflection.Register(s)

	err = s.Serve(listner)
	if err != nil {
		log.Println(err)
		return
	}

}
