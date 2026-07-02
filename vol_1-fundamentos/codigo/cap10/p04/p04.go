package main

import "fmt"

type Describer interface {
	Description() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Description() string {
	return fmt.Sprintf("Nome: %s - idade %d", p.Name, p.Age)
}

type Product struct {
	Name   string
	Weight int
	Cost   int
}

func (p Product) Description() string {
	return fmt.Sprintf("Produto: %s, %d g, %.2f R$/kg", p.Name, p.Weight, float64(p.Cost)/100.)
}

func printDescription(d Describer) {
	fmt.Println(d.Description())
}

func main() {
	a := Person{"Fred", 35}
	b := Product{"Salame", 500, 1800}

	printDescription(a)
	printDescription(b)
}
