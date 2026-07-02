package main

import "fmt"

func splitEvenOdd(values []int) ([]int, []int) {
	isEven := func(v int) bool {
		return v%2 == 0
	}

	var evens, odds []int

	for _, v := range values {
		if isEven(v) {
			evens = append(evens, v)
			continue
		}

		odds = append(odds, v)
	}

	return evens, odds
}

func main() {
	numbers := []int{3, 8, 5, 12, 7, 10}

	evens, odds := splitEvenOdd(numbers)

	fmt.Println(evens)
	fmt.Println(odds)
}
