package main

import (
	"log"
	"net/http"
)

// Define a home handler function with writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	// Use the http.NewServerMux() function to initialize a new servermux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters:  the TCP network address to listen on (in this case ":4000")
	// and the servermux we just created. If http.LsitenAndServe() return an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
