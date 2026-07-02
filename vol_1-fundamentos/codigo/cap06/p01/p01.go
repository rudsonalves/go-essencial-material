package main

import "fmt"

func minMax(numbers []int) (int, int) {
	min, max := 9999, -9999

	for _, v := range numbers {
		if v > max {
			max = v
			continue
		}
		if v < min {
			min = v
		}
	}

	return min, max
}

func main() {
	numbers := []int{15, 7, 42, 18, 9}

	min, max := minMax(numbers)

	fmt.Println("Mínimo:", min)
	fmt.Println("Máximo:", max)
}
