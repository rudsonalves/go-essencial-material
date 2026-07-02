package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value := os.Args
	if len(value) != 2 {
		fmt.Println("Passe um valor numárico!")
		return
	}

	num, err := strconv.Atoi(value[1])
	if err != nil {
		fmt.Println("O valor passado deve ser um número!")
		return
	}

	mult := 1
	for i := 1; i <= num; i++ {
		mult *= i
	}

	fmt.Printf("%d! = %d\n", num, mult)
}
