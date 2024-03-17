package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice as the response body
func home(w http.ResponseWriter, r *http.Request) {

	// Check if the request URL path exactly matches '/'
	// if not, response will be a 404 error to the client
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler
func createSnippet(w http.ResponseWriter, r *http.Request) {

	// If the method isn't POST, thw ResponseWriter will send a 405 status code and
	// he w.Write() method will be used to return "Method Not Allowed" in the response body
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return // doesn't allow the rest of the function body to run
	}

	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// Register a servemux using the http.NewServeMux to register the home function as the handler for the '/' URL pattern
	var mux = http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Use http.ListenAndServe() to start a new web server.
	log.Println("Starting a server on :4000")
	err := http.ListenAndServe(":4000", mux) // takes two parameters: TCP network address and the servemux we created
	log.Fatal(err)                           // if the function returns an error we use it to log the error and exit
}
