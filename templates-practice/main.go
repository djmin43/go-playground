package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
)

type Article struct {
	ID      int    `json:"id333"`
	Title   string `json:"title333"`
	Content string `json:"content333"`
}

var router *chi.Mux
var db *sql.DB

func main() {
	article := Article{1, "title1", "content1"}
	jsonData, err := json.Marshal(article)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

}
