package main

import (
	"html/template"
	"log"
	"net/http"
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

func getApi(w http.ResponseWriter, r *http.Request) {

}

func postApi(w http.ResponseWriter, r *http.Request) {

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
