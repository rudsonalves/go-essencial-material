package main

import "fmt"

type Discountable interface {
	Discount() float64
}

type Book struct {
	Name  string
	Autor string
	Valor int
}

func (b Book) String() string {
	return fmt.Sprintf("%s (%s), R$ %.2f", b.Name, b.Autor, float64(b.Valor)/100.)
}

func (b Book) Discount() float64 {
	return float64(b.Valor) * 0.85 / 100
}

type Eletronic struct {
	Nome        string
	Alimentacao int
	Valor       int
}

func (e Eletronic) String() string {
	return fmt.Sprintf("%s (%d V), R$ %.2f", e.Nome, e.Alimentacao, float64(e.Valor)/100)
}

func (e Eletronic) Discount() float64 {
	return float64(e.Valor) * 0.80 / 100
}

func printPromocao(d Discountable) {
	fmt.Printf("%s por R$ %.2f\n", d, d.Discount())
}

func main() {
	b := Book{"O Universo Numa Casca De Noz", "Stephen Hawking", 3590}
	e := Eletronic{"Câmera Smart Segurança Baba Eletrônica", 220, 5990}

	printPromocao(b)
	printPromocao(e)
}
