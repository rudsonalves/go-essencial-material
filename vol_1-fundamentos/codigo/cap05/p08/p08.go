package main

import "fmt"

func isEqual(matrix [3][3]int) bool {
	first := matrix[0][0]

	for i := 1; i < len(matrix); i++ {
		if first != matrix[i][i] {
			return false
		}
	}

	return true
}

func main() {
	matrix := [3][3]int{
		{5, 1, 2},
		{3, 5, 4},
		{7, 8, 5},
	}

	if isEqual(matrix) {
		fmt.Println("Elementos da diagonal são iguais")
	} else {
		fmt.Println("Elementos da diagonal são diferentes")
	}
}
