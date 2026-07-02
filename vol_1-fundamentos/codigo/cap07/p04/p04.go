package main

import "fmt"

func maxPtr(v1, v2 *int) *int {
	if *v1 > *v2 {
		return v1
	}

	return v2
}

func main() {
	a, b := 5, 12

	c := maxPtr(&a, &b)

	fmt.Println(a, b)
	fmt.Println(*c)

	*c = 111
	fmt.Println(*c)
	fmt.Println(a, b)
}
