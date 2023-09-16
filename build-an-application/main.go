package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	handler := http.HandlerFunc(server.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
