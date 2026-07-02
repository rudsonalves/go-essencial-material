package main

import "fmt"

func reverse(values *[]int) {
	for i, j := 0, len(*values)-1; i < j; i, j = i+1, j-1 {
		t := (*values)[i]
		(*values)[i] = (*values)[j]
		(*values)[j] = t
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println(numbers)
	reverse(&numbers)
	fmt.Println(numbers)
}
