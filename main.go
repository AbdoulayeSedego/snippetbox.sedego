package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {

	// Check if the current request URL path exactly matches "/" . If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}

// snippetview handler function
func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// createSnippet handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a new snippet"))
}

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
