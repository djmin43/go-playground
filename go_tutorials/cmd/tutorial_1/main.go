// look for the entry function here. 'main'
package main

import "fmt"

func main() {
	//var intArr [3]int32 = [...]int32{1, 2, 3}
	intArr1 := [3]int32{1, 2, 3}
	fmt.Println(&intArr1[1])

	// Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data.

	var someSlice = []int32{3, 4, 5, 6}
	fmt.Printf("capacity of someSlice is %v", cap(someSlice))
	someSlice = append(someSlice, 7)
	fmt.Printf("capacity of someSlice is %v", cap(someSlice))

	var intSlice3 []int32 = make([]int32, 3)
	intSlice3 = append(intSlice3, 6)
	fmt.Println(intSlice3)

	var someMap map[string]uint8 = map[string]uint8{"Adam": 23, "Sarah": 45}
	fmt.Println(someMap["Adam"])
	fmt.Println(someMap["Json"])

	// map does not guarantee the order
	for name := range someMap {
		fmt.Printf("Name: %v\n", name)
	}

	for index, value := range intSlice3 {
		fmt.Printf("someNumber: %v, %v\n", index, value)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	var i int32 = 0
	for {
		if i >= 10 {
			break
		}
		fmt.Println(i)
		i++
	}
	fmt.Println("hello")
}
