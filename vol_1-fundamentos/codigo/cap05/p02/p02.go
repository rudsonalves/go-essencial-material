package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	last := numbers[len(numbers)-1]

	for i := len(numbers) - 1; i > 0; i-- {
		numbers[i] = numbers[i-1]
	}

	numbers[0] = last

	fmt.Println(numbers)
}
