package main

import (
	"fmt"
	"os"
)

func main() {

	age := 35
	name := "mindongjoon"

	//Print
	fmt.Print("hello, \n")
	fmt.Print("world\n")

	// Println
	fmt.Println("new line")
	fmt.Println("new line again")
	fmt.Println("my age is ", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("my name is %v and my name is %v \n", age, name)
	fmt.Printf("my name is %q and my name is %q\n", age, name)
	fmt.Printf("age is of type %T\n", age)
	fmt.Printf("you scored %0.1f points\n", 225.55)

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("error")
		return
	}
	defer file.Close()

	fmt.Fprint(file, "testing Fprintf")

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my name is %v and my name is %v \n", age, name)
	fmt.Println("the saved string is: ", str)

}
