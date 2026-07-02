package main

import "fmt"

func sum(a, b *int) int {
	return *a + *b
}

func main() {
	i, j := 2, 7

	c := sum(&i, &j)

	fmt.Println(c)
}
