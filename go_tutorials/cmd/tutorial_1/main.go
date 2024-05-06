// look for the entry function here. 'main'
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		go dbCall(i)
	}

	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
}

func dbCall(i int) {
	//	Simulate DB call delay
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is: ", dbData[i])
}
