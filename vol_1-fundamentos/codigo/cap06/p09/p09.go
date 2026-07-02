package main

import "fmt"

func counter() func() int {
	var count = 0

	return func() int {
		count++
		return count
	}
}

func main() {
	c1 := counter()
	c2 := counter()

	fmt.Println("Counter 1")
	fmt.Println(c1())
	fmt.Println(c1())

	fmt.Println("\nCounter 2")
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
	fmt.Println(c2())
}
