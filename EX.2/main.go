package main

import (
	"fmt"
	"net/http"
)

// Define a map of valid usernames and passwords.
var validUsers = map[string]string{
	"john": "password123",
	"bob":  "password456",
}

// Define the authentication middleware function.
func authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the username and password from the request headers.
		username, password, ok := r.BasicAuth()

		// Check if the username and password are valid.
		if !ok || validUsers[username] != password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Please enter your username and password."`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Unauthorized\n")
			return
		}

		// If the username and password are valid, call the next handler.
		next.ServeHTTP(w, r)
	})
}

// Define the handler function for the protected endpoint.
func protectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are authenticated!\n")
}

// Define the main function.
func main() {
	// Create a new mux.
	mux := http.NewServeMux()

	// Register the protected endpoint with the authentication middleware.
	mux.Handle("/login", authenticate(http.HandlerFunc(protectedHandler)))

	// Start the server.
	fmt.Println("Server listening on port 4000...")
	http.ListenAndServe(":4000", mux)
}
