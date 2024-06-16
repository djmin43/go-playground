package main

import "fmt"

func do(m1 *map[int]int) {
	(*m1)[3] = 0
	fmt.Println("m1", *m1)
	*m1 = make(map[int]int)
	(*m1)[4] = 4
	fmt.Println("m1", *m1)
}

func main() {
	m := map[int]int{4: 1, 7: 2, 8: 3}

	fmt.Println("m", m)

	do(&m)
	do(&m)

	fmt.Println("m", m)

}
