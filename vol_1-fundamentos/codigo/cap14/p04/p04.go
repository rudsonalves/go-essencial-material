package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("Uso: go run main.go <delay_em_segundos>")
		os.Exit(1)
	}

	delay, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Argumento inválido. Tente novamente.")
		os.Exit(1)
	}

	if delay <= 0 {
		fmt.Println("O argumento deve ser um número inteiro positivo.")
		os.Exit(1)
	}

	finish := time.Now().Add(time.Duration(delay) * time.Second)
	for {
		now := time.Now()

		if now.After(finish) {
			break
		}

		fmt.Println(now.Format("15:04:05"))
		time.Sleep(time.Second)
	}

	fmt.Println("Terminou!!!")
}
