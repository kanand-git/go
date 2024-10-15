package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	mux := chi.NewRouter()

	//r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	if r.Method != http.MethodGet {
	//		return
	//	}
	//	w.Write([]byte("Hello World"))
	//})
	// mux.Use applies middlewares to all the routes using the muxer
	mux.Use(middleware.Logger, middleware.Recoverer)
	mux.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		response := map[string]string{
			"message": "Hello, JSON! (Chi)",
			"status":  "success",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	})

	//localhost:8080/user/123
	// Route Parameters
	mux.Get("/user/{name}", func(w http.ResponseWriter, r *http.Request) {
		name := chi.URLParam(r, "name")
		w.Write([]byte("Hello, " + name + "! (Chi)"))
	})

	// Query Parameters
	mux.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		firstName := r.URL.Query().Get("first_name")
		if firstName == "" {
			firstName = "Guest"
		}
		lastName := r.URL.Query().Get("lastName")
		w.Write([]byte("Hello, " + firstName + " " + lastName + "!"))
	})

	// Grouping Routes
	mux.Route("/v1", func(r chi.Router) {
		// below middleware would be applied to the current group only
		r.Use(middleware.Logger)

		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Users v1 (Chi)"))
		})
		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Posts v1 (Chi)"))
		})
	})
	mux.Route("/v2", func(r chi.Router) {
		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Users v2 (Chi)"))
		})
		r.Get("/posts", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Posts v2 (Chi)"))
		})
	})
	// using mux from chi to match the request not the http
	http.ListenAndServe(":8080", mux)

}

/*
# GET /json
curl http://localhost:8082/json

# GET /user/{name} (example with 'name' as 'John')
curl http://localhost:8082/user/John

# GET /welcome with query parameters
curl "http://localhost:8082/welcome?first_name=John&lastName=Doe"

# GET /v1/users
curl http://localhost:8082/v1/users

# GET /v1/posts
curl http://localhost:8082/v1/posts

# GET /v2/users
curl http://localhost:8082/v2/users

# GET /v2/posts
curl http://localhost:8082/v2/posts


*/
