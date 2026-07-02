package main

import "fmt"

func isPositive(v int) bool {
	return v > 0
}

func main() {
	numbers := []int{-5, 8, -2, 10, 0, 7, -1}
	evens := []int{}

	for _, n := range numbers {
		if isPositive(n) {
			evens = append(evens, n)
		}
	}

	fmt.Println("São positivos", evens)
}
