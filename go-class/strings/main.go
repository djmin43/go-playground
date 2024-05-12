package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "not enough args")
		os.Exit(-1)
	}

	old, ne := os.Args[1], os.Args[2]

	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		s := strings.Split(scan.Text(), old)

		fmt.Printf("%v", len(s))
		t := strings.Join(s, ne)

		fmt.Println(t)
	}
}
