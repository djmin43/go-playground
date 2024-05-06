// look for the entry function here. 'main'
package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type electricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up first!")
	}
}

func main() {
	var myEngine = gasEngine{mpg: 25, gallons: 33, owner: owner{name: "ff"}, int: 33}
	var yourEngine = electricEngine{kwh: 4, mpkwh: 3}
	canMakeIt(myEngine, 50)
	canMakeIt(yourEngine, 50)
}
