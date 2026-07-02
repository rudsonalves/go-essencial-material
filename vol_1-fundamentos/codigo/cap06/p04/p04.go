package main

import "fmt"

func countEvenOdd(values ...int) (int, int) {
	even, odd := 0, 0

	isEven := func(v int) bool {
		return v%2 == 0
	}

	for _, v := range values {
		if isEven(v) {
			even++
		} else {
			odd++
		}
	}

	return even, odd
}

func main() {
	numbers := []int{3, 8, 5, 12, 7, 10, 13}

	even, odd := countEvenOdd(numbers...)

	fmt.Println("São", even, "pares e", odd, "ímpares")
}
