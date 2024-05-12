package main

import "fmt"

func main() {

	t := []byte("string")

	fmt.Println(len(t), t)
	fmt.Println(t[2])
	fmt.Println(t[:2])
	fmt.Println(t[3:5])

}
