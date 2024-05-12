// look for the entry function here. 'main'
package main

import "fmt"

func main() {
	var c = make(chan int)
	go process(c)
	fmt.Println(<-c)

}

func process(c chan int) {
	c <- 123

}
