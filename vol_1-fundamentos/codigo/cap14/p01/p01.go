package main

import (
	"fmt"
)

type Person struct {
	Name   string
	Age    int
	Height float64
}

func (p Person) String() string {
	return fmt.Sprintf("Nome: %s, idade: %d, altura: %.2f", p.Name, p.Age, p.Height)
}

func main() {
	p := Person{}

	fmt.Print("Entre seu nome: ")
	fmt.Scan(&p.Name)
	fmt.Print("Entre idade altura: ")
	for {
		n, err := fmt.Scanf("%d %f", &p.Age, &p.Height)
		if err == nil && n == 2 {
			break
		}
		fmt.Print("Entre a idade e a altura separados por um espaço: ")
	}

	fmt.Println(p)
}
