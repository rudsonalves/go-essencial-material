package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	value := os.Args
	if len(value) != 3 {
		fmt.Println("Passe a base e o expoente!")
		return
	}

	base, err := strconv.Atoi(value[1])
	if err != nil {
		fmt.Println("Base deve ser um número!")
		return
	}

	expoente, err := strconv.Atoi(value[2])
	if err != nil || expoente < 0 {
		fmt.Println("Expoente deve ser um número >= 0!")
		return
	}

	mult := 1
	for i := 0; i < expoente; i++ {
		mult *= base
	}

	fmt.Printf("%dˆ%d = %d\n", base, expoente, mult)
}
