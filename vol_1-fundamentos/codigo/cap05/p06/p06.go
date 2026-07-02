package main

import "fmt"

func isSorted(vec []int) bool {

	for index := 1; index < len(vec); index++ {
		value := vec[index]
		before := vec[index-1]
		if before > value {
			return false
		}
	}

	return true
}

func main() {
	numbers := []int{3, 7, 10, 12, 15, 18}

	if isSorted(numbers) {
		fmt.Println("Vetor está ordenado")
	} else {
		fmt.Println("Vetor não está ordenado")
	}
}
