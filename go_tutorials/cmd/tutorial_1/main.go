// look for the entry function here. 'main'
package main

import "fmt"

func main() {
	// store the pointer
	// * this star is just a pointer type. new(int32) intiailazes the p value to a pointer.
	var p *int32 = new(int32)

	// this is value
	var i int32

	// ampersand is the reference to the memory address of the variable
	p = &i
	*p = 1

	// This star sign has a different duty. this sends the value of the pointer value.
	fmt.Printf("The value p points to is: %v", *p)
	fmt.Printf("\nThe value if i is: %v", i)

	var slice = []int{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[2] = 4
	fmt.Printf("\nThe value slice points to is: %v", slice)
	fmt.Printf("\nThe value slice points to is: %v", sliceCopy)

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v", result)
}

func square(thing2 *[5]float64) [5]float64 {
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2

}
