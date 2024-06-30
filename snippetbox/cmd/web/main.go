package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// http.Dir는 root를 기준으로 파일 경로를 만들어줘
	// file server는 hanlder를 리턴함
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// ex: 요청이 /static/some-file.css로 들어오게 됨
	// 그러면 /static을 없애고 나서 relative path에 some-file.css를 뽑아줌
	// ./ui/static/some-file.css가 된다.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Listening on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)
}
