package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	value := rand.IntN(20) + 1

	count := 1
	for {
		if count == value {
			break
		}

		fmt.Println(count)
		count++
	}

	fmt.Println(value)
}
