package main

import "fmt"

func main() {
	numbers := [10]int{3, 7, 5, 9, 1, 4, 8, 2, 6, 1}

	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				fmt.Printf("índice %d e %d possuem o mesmo elemento %d\n", i, j, numbers[i])
				return
			}
		}
	}

	fmt.Println("São todos diferentes")
}
