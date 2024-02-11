package main

import (
	"net/http"
)

// need to know why it has return func type here
func (app *application) routes() *http.ServeMux {

	// mux is a handler
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)

	return mux
}
