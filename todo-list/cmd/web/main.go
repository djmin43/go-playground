package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("hello")

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	http.ListenAndServe(":4000", mux)

}

func home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "hello world")
}
