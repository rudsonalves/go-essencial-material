package main

import "fmt"

type Shape interface {
	Area() float64
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.1415 * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Println("Área:", s.Area())
}

func main() {
	radius := 1.5
	c := Circle{radius}
	s := Square{2 * radius}

	printArea(c)
	printArea(s)
}
