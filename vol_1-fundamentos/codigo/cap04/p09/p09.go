package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		num1, num2         int
		val1, val2, signal string
	)

	fmt.Print("Entre uamexpressão como 2 + 5: ")
	fmt.Scanf("%s %s %s", &val1, &signal, &val2)

	num1, err1 := strconv.Atoi(val1)
	num2, err2 := strconv.Atoi(val2)

	if err1 != nil || err2 != nil {
		fmt.Println("Entrada incorreta!")
		return
	}

	switch signal {
	case "+":
		fmt.Printf("%d + %d = %d\n", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d\n", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d\n", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d\n", num1, num2, num1/num2)
	case "%":
		fmt.Printf("%d %% %d = %d\n", num1, num2, num1%num2)
	default:
		fmt.Println("Sinal errado!")
	}
}
