package main

import "fmt"

func appendNew(set []int, value int) []int {
	for _, v := range set {
		if v == value {
			return set
		}
	}

	return append(set, value)
}

func main() {
	numbers := []int{3, 7, 3, 2, 5, 7, 9, 3}

	repeated := []int{}

	for i, v := range numbers {
		for j := i + 1; j < len(numbers); j++ {
			if v == numbers[j] {
				repeated = appendNew(repeated, v)
			}
		}
	}

	fmt.Println(repeated)
}
