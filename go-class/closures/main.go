package main

import "fmt"

func main() {
	f, g := fib(), fib()

	fmt.Println(f(), f(), f(), f(), f())
	fmt.Println(g(), g(), g(), g(), g())
}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
