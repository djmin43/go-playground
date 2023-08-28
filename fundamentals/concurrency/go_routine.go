package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main 함수 시작", time.Now())

	done := make(chan bool)
	go long(done)
	go short(done)

	<-done
	<-done

	time.Sleep(5 * time.Second) // 5초 대기
	fmt.Println("main 함수 종료", time.Now())
}

func long(done chan bool) {
	fmt.Println("long 함수 시작", time.Now())
	time.Sleep(3 * time.Second)
	fmt.Println("long 함수 종료", time.Now())
	done <- true
}

func short(done chan bool) {
	fmt.Println("short 함수 시작", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("short 함수 종료", time.Now())
	done <- true
}
