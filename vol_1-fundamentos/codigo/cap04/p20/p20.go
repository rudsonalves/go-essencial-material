package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40}

	sum := 0
	for _, n := range numbers {
		sum += n
	}

	medium := sum / len(numbers)
	fmt.Println(medium)
}
