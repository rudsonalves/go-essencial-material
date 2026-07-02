package main

import (
	"fmt"
	"strconv"
)

func main() {
	var value string

	fmt.Print("Entre um número: ")
	fmt.Scanln(&value)

	number, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Erro ao converter o número:", err)
		return
	}
	fmt.Println("Número digitado:", number)

	if number == 0 {
		fmt.Println("O número é zero.")
	} else if number > 0 {
		fmt.Println("O número é positivo.")
	} else {
		fmt.Println("O número é negativo.")
	}
}
