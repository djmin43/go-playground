package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":3000", "HTTP network address")

	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", todoView)
	mux.HandleFunc("/create", todoCreate)

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Server running on %s", *addr)
	err := http.ListenAndServe(*addr, mux)

	if err != nil {
		log.Fatal(err)
	}

}
