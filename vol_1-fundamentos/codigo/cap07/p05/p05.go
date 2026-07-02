package main

import "fmt"

func double(x *int) {
	*x *= 2
}

func main() {
	number := 10

	double(&number)

	fmt.Println(number)
}
