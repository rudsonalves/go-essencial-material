package main

import (
	"fmt"
	"strconv"
)

func str2Int(value string) (int, error) {
	number, err := strconv.Atoi(value)

	return number, err
}

func main() {
	var (
		num1, num2 int
		val1, val2 string
	)

	fmt.Println("Entrecom dois números separados por um espaço: ")
	fmt.Scanf("%s %s", &val1, &val2)

	num1, err := str2Int(val1)
	if err != nil {
		fmt.Printf("Ocorreu um erro: %s\n", err)
		return
	}

	num2, err = str2Int(val2)
	if err != nil {
		fmt.Printf("Ocorreu um erro: %s\n", err)
		return
	}

	if num1 == num2 {
		fmt.Println("Os números devem ser difereentes!")
		return
	}

	if num1 > num2 {
		fmt.Printf("%d é maior que %d\n", num1, num2)
	} else {
		fmt.Printf("%d é menor que %d\n", num1, num2)
	}
}
