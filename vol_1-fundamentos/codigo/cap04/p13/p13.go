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

	for i := 1; i <= 10; i++ {
		fmt.Printf("%2d * %2d = %2d\n", num, i, num*i)
	}
}
