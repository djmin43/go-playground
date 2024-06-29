package main

import (
	"log"
	"net/http"
)

func main() {

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

	log.Println("Listening on 4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
