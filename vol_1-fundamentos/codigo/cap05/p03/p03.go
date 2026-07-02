package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{10, 20, 30, 40}

	sum := []int{0, 0, 0, 0}

	for i := 0; i < len(a); i++ {
		sum[i] = a[i] + b[i]
	}

	fmt.Println("Soma:", sum)
}
