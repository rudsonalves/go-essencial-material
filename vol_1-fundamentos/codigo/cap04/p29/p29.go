package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Deve passar um argumento numérico!")
		return
	}

	num, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Entre um número inteiro!")
		return
	}

	for i := 2; i < num/2; i++ {
		if num%i == 0 {
			fmt.Printf("%d é divisível por %d. Não é primo!\n", num, i)
			return
		}
	}

	fmt.Printf("%d é primo!\n", num)
}
