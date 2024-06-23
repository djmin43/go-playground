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
