package main

import "fmt"

func main() {

	n, err := fmt.Println("Hello, playground")

	if _, err := fmt.Println(n); err != nil {
		fmt.Println(err)
	}
}
