package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	first := numbers[0]
	newSlice := numbers[1:][:]
	newNumbers := append([]int{}, newSlice...)
	newNumbers = append(newNumbers, first)

	numbers = newNumbers

	fmt.Println(numbers)
}
