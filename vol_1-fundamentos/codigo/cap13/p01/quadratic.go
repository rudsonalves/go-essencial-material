package main

import (
	"fmt"
	"os"
	"quadratic/solver"
	"strconv"
)

func formatValue(v float64) string {
	switch {
	case v < 0:
		return fmt.Sprintf(" - %.2f", -v)
	case v == 0:
		return ""
	default:
		return fmt.Sprintf(" + %.2f", v)
	}
}

func printEquation(a, b, c float64) {
	if b == 0 {
		fmt.Printf("%.2fx²%s = 0\n", a, formatValue(c))
	} else {
		fmt.Printf("%.2fx²%sx%s = 0\n", a, formatValue(b), formatValue(c))
	}
}

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Use: quadratic 1 -3 2\nCom aomenos um argumento.")
		fmt.Printf("onde o polinômio a resolver é: 1x² + (-3)x + 2 = 0")
		os.Exit(1)
	}

	coefs := [3]float64{}
	for index, arg := range args {
		coef, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Printf("Coeficiente \"%s\" não é um número\n", arg)
			os.Exit(1)
		}
		coefs[index] = coef
	}

	x1, x2, err := solver.Solver(coefs[0], coefs[1], coefs[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	printEquation(coefs[0], coefs[1], coefs[2])
	fmt.Printf("Raízes: x1 = %.2f; x2 = %.2f\n", x1, x2)
}
