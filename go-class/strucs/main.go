package main

import (
	"fmt"
	"time"
)

type Employee struct {
	Name   string
	Number int
	Boss   *Employee
	Hired  time.Time
}

func main() {

	c := map[string]*Employee{}

	c["Lamine"] = &Employee{"Lamine", 2, nil, time.Now()}

	c["Matt"] = &Employee{
		Name:   "Matt",
		Number: 1,
		Hired:  time.Now(),
		Boss:   c["Lamine"],
	}

	fmt.Printf("%T %+[1]v\n", c["Lamine"])
	fmt.Printf("%T %+[1]v\n", c["Matt"])
}
