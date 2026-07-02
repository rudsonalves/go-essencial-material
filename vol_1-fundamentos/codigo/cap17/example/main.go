package main

import "fmt"

func Print(value any) {
	switch v := value.(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("Inteiro: %d\n", v)

	case float32, float64:
		fmt.Printf("Float: %.2f\n", v)

	case string:
		fmt.Printf("Texto: %s\n", v)

	default:
		fmt.Printf("Tipo desconhecido: %T\n", v)
	}
}

func main() {
	Print(10)
	Print(3.14)
	Print("Go")
	Print(true)
}
