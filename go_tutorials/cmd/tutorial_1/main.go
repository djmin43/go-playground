// look for the entry function here. 'main'
package main

import (
	"fmt"
	"strings"
)

func main() {

	var myString = []rune("résumé")
	for i, v := range myString {
		fmt.Println(i, v)
	}

	// String is just byte. rune is the unicode point number
	// rune i just alias for int32
	var myRune = 'a'
	fmt.Printf("\n my Rune = %v, %T", myRune, myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var strBuilder strings.Builder
	for i := range strSlice {
		strBuilder.WriteString(strSlice[i])
	}
	var catStr = strBuilder.String()
	fmt.Println(catStr)
}
