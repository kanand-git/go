package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"small-app/models"
	"strconv"
)

type Controller struct {
	// s would be nil by default
	s *models.Service // handler depends on models.Service
}

func NewController(s *models.Service) (*Controller, error) {
	if s == nil {
		return nil, errors.New("service is nil")
	}
	return &Controller{s: s}, nil
}

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	// this line set your  ContentType as json
	w.Header().Set("Content-Type", "application/json")

	//fetching the variable from query // ?user_id =123
	userIdString := r.URL.Query().Get("user_id")

	//converting it to make sure it is a valid uint64
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		log.Println(err) // this is internal logging

		appErr := map[string]any{"Message": "Please provide a valid user id"}
		log.Println("sent the resp to end user", appErr)

		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(appErr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return

	}

	// fetching the user from db
	//u, err := models.FetchUser(userId)

	u, err := c.s.FetchUser(userId)

	if err != nil {
		// internal app logging
		log.Println(err)

		// creating a generic message for the end user, don't reveal internal details
		appErr := map[string]any{"Message": "User not found"}
		log.Println("sent the resp to end user", appErr)

		// setting the status code and sending the responses
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(appErr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
