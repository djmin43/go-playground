package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// http.Dir는 root를 기준으로 파일 경로를 만들어줘
	// file server는 hanlder를 리턴함
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// ex: 요청이 /static/some-file.css로 들어오게 됨
	// 그러면 /static을 없애고 나서 relative path에 some-file.css를 뽑아줌
	// ./ui/static/some-file.css가 된다.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Listening on %s", *addr)
	err := srv.ListenAndServe()

	errorLog.Fatal(err)
}
