package main

import "fmt"

func main() {
	a, b := 17, 345
	c, d := 1.2, 3.45

	fmt.Printf("|%-9d|%9d|\n", a, b)
	fmt.Printf("|%9f|%9.2f|\n", c, d)
}
