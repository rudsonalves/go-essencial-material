package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func drawNumber() int {
	return rand.IntN(99)
}

func main() {
	fmt.Println("Un número de 0 a 99 foi sorteado. Descubra o número com o menor número de tentativas.")

	count := 0
	number := drawNumber()
	for {
		count++

		var input string
		fmt.Printf("Entre com um número [%d]: ", count)
		n, err := fmt.Scan(&input)
		if err != nil || n == 0 {
			fmt.Println()
			continue
		}

		value, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println()
			continue
		}

		if value == number {
			break
		}

		if value > number {
			fmt.Printf("\n%d é maior que o número sorteado!\n", value)
		} else {
			fmt.Printf("\n%d é menor que o número sorteado!\n", value)
		}
	}

	fmt.Printf("Acertou em %d tentativas o numero sorteado %d.\n", count, number)
}
