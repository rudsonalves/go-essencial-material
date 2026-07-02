package domain

import "fmt"

type Product struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func NewProduct(name string, price int) Product {
	return Product{
		Name:  name,
		Price: price,
	}
}

func (p Product) String() string {
	return fmt.Sprintf("%-12s - R$%.2f", p.Name, float64(p.Price)/100)
}
