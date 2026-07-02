package main

import "fmt"

func main() {
	matrix := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	transposta := [3][3]int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			transposta[j][i] = matrix[i][j]
		}
	}

	fmt.Println(transposta)
}
