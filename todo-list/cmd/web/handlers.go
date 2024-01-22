package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "Get Method Required", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "hello world")
}
