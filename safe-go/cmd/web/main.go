package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/login_user", getLoginUser)
	mux.HandleFunc("/log", postLog)
	mux.HandleFunc("/attacker_example", getAttackerExample)

	log.Printf("Starting server on %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatalln(err)

}
