package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getApi(w, r)
	case http.MethodPost:
		postApi(w, r)
	}

}

type Message struct {
	Message string `json:"message"`
}

func getApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4000")
	fmt.Printf(w.Header().Get("Access-Control-Allow-Origin"))
	err := json.NewEncoder(w).Encode(&Message{"hello"})
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func postApi(w http.ResponseWriter, r *http.Request) {

	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prettyJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("X-Timestamp", strconv.Itoa(int(time.Now().Unix())))
	lang := r.Header.Get("x-lang")
	fmt.Printf("%s", lang)
	fmt.Println("else")
	fmt.Fprintf(w, "%s\n", string(prettyJSON))
	fmt.Printf("%s\n", string(prettyJSON))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	getHome(w, r)
}

func getHome(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/pages/index.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func getLoginUser(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/pages/login_user.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func getAttackerExample(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/pages/attacker_example.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func postLog(w http.ResponseWriter, r *http.Request) {

	var m Message
	if err := json.NewDecoder(r.Body).Decode(&m); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	prettyJSON, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%s\n", string(prettyJSON))
}
