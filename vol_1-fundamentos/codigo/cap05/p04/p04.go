package main

import "fmt"

func isDiagonal(matrix [3][3]int) bool {
	for i, row := range matrix {
		for j, e := range row {
			if i != j && e != 0 {
				return false
			}
		}
	}

	return true
}

func main() {
	matrix := [3][3]int{
		{2, 0, 0},
		{0, 5, 0},
		{0, 0, 8},
	}

	if isDiagonal(matrix) {
		fmt.Println("É uma matriz diagonal")
	} else {
		fmt.Println("Não é uma matriz diagonal")
	}
}
