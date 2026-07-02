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
		num1, num2, num3, great int
		val1, val2, val3        string
	)

	fmt.Print("Entre três números: ")
	fmt.Scanf("%s %s %s", &val1, &val2, &val3)

	num1, err1 := str2Int(val1)
	num2, err2 := str2Int(val2)
	num3, err3 := str2Int(val3)

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println("Deu err na entrada de um dos números!")
		return
	}

	great = num1

	if num2 > great {
		great = num2
	}

	if num3 > great {
		great = num3
	}

	fmt.Printf("O maior número em [%d, %d, %d] é %d\n", num1, num2, num3, great)
}
