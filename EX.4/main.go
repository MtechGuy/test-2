package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	// create a new mux to handle incoming requests
	mux := http.NewServeMux()

	// add middleware to check Content-Type header
	mux.HandleFunc("/", contentTypeMiddleware(handleRequest))

	// start the server on port 4000
	fmt.Println("Server listening on port 4000...")
	http.ListenAndServe(":4000", mux)
}

// define a struct to hold the JSON payload
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// handleRequest processes incoming requests
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// decode the JSON payload into a User struct
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// print the user's name and email to the console
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)

	// write a response back to the client
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", user.Name)
}

// contentTypeMiddleware checks for the Content-Type header and
// stops the request if it's not application/json
func contentTypeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		if contentType != "application/json" {
			http.Error(w, "Content-Type header must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		next.ServeHTTP(w, r)
	}
}
