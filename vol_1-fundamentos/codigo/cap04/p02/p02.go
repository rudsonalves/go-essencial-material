package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		age   int
		value string
	)

	fmt.Print("Entre a sua idade: ")
	fmt.Scanln(&value)

	age, err := strconv.Atoi(value)
	if err != nil {
		fmt.Printf("\"%s\" não é um número.\n", value)
		return
	}

	if age < 18 {
		fmt.Println("Menor de idade!")
	} else {
		fmt.Println("Maior de idade!")
	}
}
