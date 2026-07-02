package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	values := os.Args

	if len(values) <= 1 {
		fmt.Println("Passe algum valor numérico na linha de commando!")
		return
	}

	num, err := strconv.Atoi(values[1])
	if err != nil {
		fmt.Println("Valor passado deve ser um número!")
		return
	}

	var sum int
	for i := 1; i < num+1; i++ {
		sum += i
	}

	fmt.Println(sum)
}
