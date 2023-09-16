package main

import (
	"net/http"
)

func home(w http.ResponseWriter, r *http.ResponseWriter) {
	w.Write([]byte("Hello from snippetbox"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

}
