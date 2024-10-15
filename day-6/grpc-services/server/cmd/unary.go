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
		return nil, status.Error(codes.Internal, "invalid user")
	}

	var user models.User
	user.Name = nu.Name
	user.Email = nu.Email
	user.Roles = nu.Roles
	v := validator.New() // Creating a new validator instance
	err := v.Struct(user)

	if err != nil {
		// If validation fails, return an error message with the error status.
		return nil, status.Error(codes.Internal, "please provide required fields in correct format")
	}

	fmt.Println(user)
	return &proto.SignupResponse{Result: user.Email + "created"}, nil

}
