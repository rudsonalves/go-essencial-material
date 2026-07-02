package main

import "fmt"

func zero(p *int) {
	*p = 0
}

func main() {
	a := 12

	fmt.Println(a)
	zero(&a)
	fmt.Println(a)
}
