package main

import "fmt"

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	mainDia := 0
	secDia := 0

	l := len(matrix)
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		mainDia += matrix[i][i]
		secDia += matrix[i][j]
	}

	fmt.Println("Diagonal principal:", mainDia)
	fmt.Println("Diagonal secundária:", secDia)
}
