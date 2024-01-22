package main

import "net/http"

// need to know why it has return func type here
func (app *application) routes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)

	//Not sure how to add file server..

	return mux
}
