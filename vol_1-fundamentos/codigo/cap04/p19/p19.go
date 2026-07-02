package main

import "fmt"

func main() {
	numbers := []int{15, 7, 42, 18, 9}

	maior := numbers[0]
	for _, n := range numbers {
		if n > maior {
			maior = n
		}
	}

	fmt.Printf("O maior é %d\n", maior)
}
