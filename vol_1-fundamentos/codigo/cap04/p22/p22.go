package main

import "fmt"

func isEven(value int) bool {
	return value%2 == 0
}

func main() {
	numbers := []int{3, 8, 5, 12, 7, 10}

	count := 0
	for _, n := range numbers {
		if isEven(n) {
			count++
		}
	}

	fmt.Println("São pares ", count)
}
