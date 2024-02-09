package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatalln(err)
}

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hello"))

}
