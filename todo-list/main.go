package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", todoView)
	mux.HandleFunc("/create", todoCreate)

	err := http.ListenAndServe(":3000", mux)

	log.Println("Server running on :3000")

	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("This is todo list"))
}

func todoView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("todo items"))
}

func todoCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create todo item"))
}
