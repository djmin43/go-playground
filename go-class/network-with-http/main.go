package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

var form = `
<h1>Todo #{{.ID}}</h1>
<div>{{ printf "User %d" .UserID }}</div>
<div>{{ printf "%s (completed: %t)" .Title .Completed }}</div>

`

type todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func handler(w http.ResponseWriter, r *http.Request) {
	const base = "https://jsonplaceholder.typicode.com/"
	resp, err := http.Get(base + r.URL.Path[1:])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var item todo

	if err = json.NewDecoder(resp.Body).Decode(&item); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl := template.New("mine")

	tmpl.Parse(form)
	tmpl.Execute(w, item)

}
