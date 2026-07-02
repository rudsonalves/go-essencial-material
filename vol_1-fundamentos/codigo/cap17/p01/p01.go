package main

import "fmt"

func Describe(value any) string {
	switch v := value.(type) {
	case string:
		return fmt.Sprintf("Texto: %s", v)
	case int:
		return fmt.Sprintf("Inteiro: %d", v)
	case float64:
		return fmt.Sprintf("Decimal: %.2f", v)
	case bool:
		return fmt.Sprintf("Booleano: %v", v)
	case nil:
		return "Valor nulo"
	default:
		return fmt.Sprintf("Tipo desconhecido: %T", v)
	}
}

func main() {
	fmt.Println(Describe(12))
	fmt.Println(Describe("Home"))
	fmt.Println(Describe(1.22))
	fmt.Println(Describe(true))
	fmt.Println(Describe([]int{}))
	fmt.Println(Describe(nil))

}
