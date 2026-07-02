package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		value string
		nota  float64
	)

	fmt.Print("Digite a nota: ")
	fmt.Scan(&value)

	nota, err := strconv.ParseFloat(value, 64)
	if err != nil {
		fmt.Println("Erro na entrada da nota!")
		return
	}

	switch {
	case nota <= 6:
		fmt.Println("Reprovado")
	case nota <= 7.9:
		fmt.Println("Recuperação")
	case nota <= 10.0:
		fmt.Println("Aprovado")
	default:
		fmt.Println("Nota não pode ser superior a 10!")
	}
}
