package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

func (p Product) String() string {
	return fmt.Sprintf("%s: R$ %.2f (Desconto %.0f%%)", p.Name, p.Price, p.Discount*100)
}

func (p Product) Advertisement() string {
	return fmt.Sprintf("%s: de R$ %.2f por R$ %.2f", p.Name, p.Price, p.FinalPrice())
}

func (p Product) FinalPrice() float64 {
	return p.Price * (1 - p.Discount)
}

func (p *Product) ApplyDiscount(discount float64) {
	if discount >= 0 && discount < 1 {
		p.Discount = discount
	}
}

func main() {
	p := Product{"Macarrão Italia 500g", 100, .15}

	fmt.Println(p)
	fmt.Println(p.Advertisement())

	p.ApplyDiscount(0.50)
	fmt.Println(p)
	fmt.Println(p.Advertisement())
}
