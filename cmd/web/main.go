package main

import (
	"log"
	"net/http"
)

func main() {

	// the http.NewServeMux() function to initialize a new servemux, then
	mux := http.NewServeMux()

	// register the home function as the handler for the "/" URL pattern.
	mux.HandleFunc("/", home)

	// Register the handler function with corresponding URL pattern
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	//port for the web server to listen to
	port := ":4000"

	log.Println("Starting web server on port ", port)
	//http.ListenAndServe() function start a new web server. take in
	// two parameters: the TCP network address to listen on and the servemux
	err := http.ListenAndServe(port, mux)

	//If http.ListenAndServe() returns an error
	// we use the log.Fatal() function to log the error message and exit.
	//any error returned by http.ListenAndServe() is always non-nil.
	log.Fatal(err)

}
