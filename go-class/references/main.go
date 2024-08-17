package main

import "fmt"

func main() {

	items := [][2]byte{{1, 2}, {3, 4}, {5, 6}}

	a := [][]byte{}

	for _, item := range items {
		i := make([]byte, len(item))

		copy(i, item[:]) // make unique
		a = append(a, i)
	}

	fmt.Println(items)
	fmt.Println(a)

}
