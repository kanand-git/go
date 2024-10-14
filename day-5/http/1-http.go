package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	// register endpoints
	http.HandleFunc("/home", home)
	http.HandleFunc("/user", FindUser)

	// run the server // it would run, until someone manually kills it
	http.ListenAndServe(":8080", nil)
	// mux // mux matches request to handler functions
	// http has a DefaultServeMux mux, which can match request to specific endpoints
	// in ListenAndServe if we pass the handler value as nil, by default it would use http.DefaultServeMux
}
func home(w http.ResponseWriter, r *http.Request) {
	//w http.ResponseWriter, is used to write resp to the client
	// http.Request// anything user send us would be in the request struct
	w.Write([]byte("Hello World"))
	//fmt.Fprintln(w, "Hello World")
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	// manually set the header to json
	w.Header().Set("Content-Type", "application/json")

	// fields must be exported , if a struct is sent to be as json
	var user struct {
		Name         string `json:"first_name"` // field level tag
		Password     string `json:"-"`          // - is to ignore the value in json output
		PasswordHash string `json:"password_hash"`
		Marks        []int  `json:"marks"`
	}
	user.Name = "John"
	user.Password = `<PASSWORD>`
	user.PasswordHash = `<PASSWORD Hash>`
	user.Marks = []int{100, 90, 80}

	//b, err := json.Marshal(user) // Marshal stores the json output in the memory
	//w.Write(b)

	// NewEncoder can directly write json to the writer
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		// signal text based error to the client
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return // don't forget to return
	}
	// setting status code
	w.WriteHeader(http.StatusOK)
}
