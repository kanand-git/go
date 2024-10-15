package models

import (
	"errors"
	"fmt"
)

type Service struct {
	db string
}

func NewService(db string) *Service {
	return &Service{
		db: db,
	}
}

var users = map[uint64]User{
	123: {
		FName: "Bob",
		LName: "abc",
		Email: "bob@email.com",
	},
}

var ErrUserNotFound = errors.New("user not found with the id ")

// FetchUser fetches the data from the map
func (s *Service) FetchUser(id uint64) (User, error) {
	u, ok := users[id]
	if !ok {

		return User{}, ErrUserNotFound
	}
	fmt.Println("fetched from", s.db)

	return u, nil
}
