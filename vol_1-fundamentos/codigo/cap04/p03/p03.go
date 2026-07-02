package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		value  string
		number int
	)

	fmt.Print("Entre com um número: ")
	fmt.Scanln(&value)

	number, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("'%s' não é um número!\n", value)
		return
	}

	par := "impar"
	if number%2 == 0 {
		par = "par"
	}

	fmt.Printf("%d é %s\n", number, par)
}
