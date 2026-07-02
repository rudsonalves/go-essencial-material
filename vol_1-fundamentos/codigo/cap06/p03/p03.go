package main

import "fmt"

func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	div, ok := divide(10, 2)
	fmt.Println("10/2:", div, ok)
	div, ok = divide(10, 0)
	fmt.Println("10/0:", div, ok)
}
