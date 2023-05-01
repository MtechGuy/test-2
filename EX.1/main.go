package main

import (
	"log"
	"net/http"
)

// write middleware
func middlewareA(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//this is executed on the way down to the handler
		log.Println("Executing middleware A")
		next.ServeHTTP(w, r)
		//this is executed on the way up to the client
		log.Println("Exectuing middleware A again")
	})
}

func middlewareB(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//this is executed on the way down to the handler
		log.Println("Executing middleware B")
		if r.URL.Path == "/cherry" {
			return
		}
		next.ServeHTTP(w, r)
		//this is executed on the way up to the client
		log.Println("Exectuing middleware B again")
	})
}

//create a handler function
func ourHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Executing the handler...")
	w.Write([]byte("CARROTS"))
}

func main() {
	mux := http.NewServeMux()                                                    //multiplexer
	mux.Handle("/check", middlewareA(middlewareB(http.HandlerFunc(ourHandler)))) //key endpoint/url and the value

	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
