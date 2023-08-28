package _select

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	fmt.Printf("%q", aDuration)
	fmt.Printf("%q", bDuration)

	if aDuration < bDuration {
		return a
	}
	return b
}
