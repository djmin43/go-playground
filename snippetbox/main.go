package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Define a home handler function with writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	// Check if the current request URL path exactly matches "/". If it doesn't, use
	// the http.NotFound() function to send a 404 response to the client.
	// Importantly, we then return from the handler. If we don't return the handler
	// would keep executing and also write the "Hello from SnippetBox" message.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.Header()["Date"] = nil
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}

func test(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed!!!"))
		return
	}
	w.Write([]byte("this is some test"))
}

func main() {
	// Use the http.NewServerMux() function to initialize a new servermux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	mux.HandleFunc("/test", test)

	// Use the http.ListenAndServe() function to start a new web server. We pass in
	// two parameters:  the TCP network address to listen on (in this case ":4000")
	// and the servermux we just created. If http.LsitenAndServe() return an error
	// we use the log.Fatal() function to log the error message and exit. Note
	// that any error returned by http.ListenAndServe() is always non-nil
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
