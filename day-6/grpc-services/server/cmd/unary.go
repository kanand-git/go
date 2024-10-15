package main

import (
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"server/gen/proto"
	"server/models"
)

func (u userService) Signup(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	nu := req.GetUser()
	if nu == nil {
		// use status.Error or Errorf to send error messages on grpc services
		return nil, status.Error(codes.Internal, "invalid user")
	}

	var user models.User
	user.Name = nu.User
	user.Email = nu.Email
	user.Roles = nu.Roles
	v := validator.New() // Creating a new validator instance

	// validating struct against the field tags specified by go validator package
	err := v.Struct(user)

	if err != nil {
		// If validation fails, return an error message with the error status.
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")
	}

	fmt.Println(user)

	// sending the success resp
	return &proto.SignupResponse{Result: user.Email + "created"}, nil

}
