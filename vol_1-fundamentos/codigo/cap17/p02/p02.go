package main

import "fmt"

func Label(value any) string {
	if s, ok := value.(fmt.Stringer); ok {
		return s.String()
	}

	return fmt.Sprintf("%v", value)
}

type Product struct {
	Name  string
	Price float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s por R$ %.2f", p.Name, p.Price)
}

func main() {
	p := Product{"Detergente Ype", 3.46}

	fmt.Println(Label(p))
	fmt.Println(Label("Casa"))
	fmt.Println(Label(12))
	fmt.Println(Label(nil))
}
