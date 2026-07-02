package main

import "fmt"

func avarage(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}

	return sum / len(values)
}

func main() {
	fmt.Println("Média:", avarage(10, 20, 30, 40))
}
