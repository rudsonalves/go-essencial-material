package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println(numbers)
}
