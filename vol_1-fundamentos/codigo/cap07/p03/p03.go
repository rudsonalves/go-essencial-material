package main

import "fmt"

func increment(value *int) {
	*value++
}

func main() {
	a := 5

	fmt.Println(a)
	increment(&a)
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)
	increment(&a)
	fmt.Println(a)
}
