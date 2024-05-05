// look for the entry function here. 'main'
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 32767
	fmt.Println(intNum)

	var myString string = "Hello world"
	fmt.Println(myString)

	// this is length of byte. go uses UTF-8.
	fmt.Println(len("test"))

	// rune is data type representing a character
	fmt.Println(utf8.RuneCountInString("f"))

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBoolean bool = true

	fmt.Println(myBoolean)

	// can't change. no initial value
	const myCont string = "const value"

}
