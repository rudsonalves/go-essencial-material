package main

import "fmt"

func minMax(values []int, min, max *int) {
	for _, v := range values {
		if *min > v {
			*min = v
			continue
		}

		if *max < v {
			*max = v
		}
	}
}

func main() {
	numbers := []int{15, 7, 42, 18, 9}

	a, b := 99, -1

	minMax(numbers, &a, &b)

	fmt.Println(a, b)
	fmt.Println(numbers)
}
