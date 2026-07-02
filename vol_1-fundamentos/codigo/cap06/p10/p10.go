package main

import "fmt"

func memoizeSquare() func(v int) int {
	mem := map[int]int{}

	return func(v int) int {
		r, ok := mem[v]

		if !ok {
			r = v * v
			mem[v] = r
			// fmt.Println(".")
		}
		return r

	}
}

func main() {
	sq := memoizeSquare()

	fmt.Println(sq(2))
	fmt.Println(sq(3))
	fmt.Println(sq(4))

	fmt.Println(sq(2))
	fmt.Println(sq(3))
}
