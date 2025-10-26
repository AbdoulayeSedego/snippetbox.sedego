package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	// The fmt.Fprintf() function interpolate the id value with the response
	// and write it to the http.ResponseWriter.
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

	//w.Write([]byte("Display a specific snippet..."))
	// set up content-Type menually
	// w.Header().Set("Content-Type", "application/json")
	// w.Write([]byte(`{"Display a specific snippet..."}`))
}

// createSnippet handler function
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	//r.Method check whether the request is using POST or not.
	// If it's not, use the w.WriteHeader() method to send a 405 status
	// code and the w.Write() method to write a "Method Not Allowed"
	// response body.
	if r.Method != http.MethodPost {
		//tell the user which Method is allowed
		w.Header().Set("Allow", http.MethodPost)
		//Surpress the Date from the header
		// w.Header()["Date"] = nil
		// w.WriteHeader(405)
		// w.Write([]byte("Method Not Allowed"))
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)

		return
	}

	w.Write([]byte("Create a new snippet"))
}
