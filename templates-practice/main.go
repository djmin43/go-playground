package main

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	_ "github.com/mattn/go-sqlite3"
	"html/template"
)

type Article struct {
	ID      int           `json:"id"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

var router *chi.Mux
var db *sql.DB

func main() {

}

func dbGetAllArticles() ([]*Article, error) {
	query, err := db.Prepare("select * from articles")
	defer query.Close()

	if err != nil {
		return nil, err
	}

	result, err := query.Query()

	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)
	for result.Next() {
		data := new(Article)
		err := result.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		if err != nil {
			return nil, err
		}
		articles = append(articles, data)

		return articles, nil
	}
}

func dbCreateArticle(article *Article) error {
	query, err := db.Prepare("insert into articles(title, content) values (?, ?)")
	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec(article.Title, article.Content)

	if err != nil {
		return err
	}

	return nil
}

func connect() (*sql.DB, error) {
	var err error
	db, err = sql.Open("sqlite3", "./data.sqlite")

	if err != nil {
		return nil, err
	}

	sqlStmt := `
    create table if not exists articles (id integer not null primary key autoincrement, title text, content text);
    `

	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	return db, nil
}
