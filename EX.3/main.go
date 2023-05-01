package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Call the next handler in the chain.
		next.ServeHTTP(w, r)

		// Log the request.
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func main() {
	// Create a new HTTP server.
	srv := &http.Server{
		Addr: ":4000",
	}

	// Create a simple handler that just returns "Hello, World!".
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, and welcome to Logging Middleware")
	})

	// Wrap the handler with the logging middleware.
	loggedHandler := LoggingMiddleware(handler)

	// Set the server's handler to the logged handler.
	srv.Handler = loggedHandler

	// Start the server.
	fmt.Println("Server listening on port 4000...")
	log.Fatal(srv.ListenAndServe())
}
