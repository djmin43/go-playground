package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
}

func main() {

	addr := flag.String("addr", ":3000", "HTTP network address")

	dsn := flag.String("dsn", "root:0000@/todolist?parseTime=true", "MySQL data source name")

	flag.Parse()

	db, err := openDB(*dsn)

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	infoLog.Printf("Listening %s", *addr)
	err = srv.ListenAndServe()

	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
