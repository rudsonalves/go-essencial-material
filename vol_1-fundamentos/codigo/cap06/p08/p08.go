package main

import "fmt"

func filter(values []int, op func(int) bool) []int {
	filtered := []int{}

	for _, v := range values {
		if op(v) {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func main() {
	numbers := []int{3, 8, 5, 12, 7, 10}

	isEven := func(v int) bool { return v%2 == 0 }
	isOdd := func(v int) bool { return v%2 != 0 }

	evens := filter(numbers, isEven)
	odds := filter(numbers, isOdd)

	fmt.Println("Pares:", evens)
	fmt.Println("Ímpares:", odds)
}
