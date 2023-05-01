package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create a new mux for routing
	mux := http.NewServeMux()

	// add route handlers
	mux.HandleFunc("/hello", helloHandler)

	// wrap the mux with the error handling middleware
	fmt.Println("Server listening on port 4000...")
	http.ListenAndServe(":4000", handleErrors(mux))
}

// handleErrors is a middleware that wraps an http.Handler and adds error handling
func handleErrors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}()
		// call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// helloHandler returns a greeting
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// simulate a panic
	if r.Method != http.MethodGet {
		panic(fmt.Sprintf("Invalid HTTP method %s. Only GET requests are allowed.", r.Method))
	}

	fmt.Fprintln(w, "Hello, and welcome to Error Handling Middleware!")
}
