package main

import "fmt"

func apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func main() {
	sum := func(a, b int) int { return a + b }
	sub := func(a, b int) int { return a - b }
	mul := func(a, b int) int { return a * b }

	fmt.Println(apply(10, 20, sum))
	fmt.Println(apply(10, 20, sub))
	fmt.Println(apply(10, 20, mul))
}
