package main

import "fmt"

func main() {
	const max, base int = 1000, 17

	number := max
	for {
		if number%base == 0 {
			break
		}

		number++
	}

	fmt.Println(number)
}
