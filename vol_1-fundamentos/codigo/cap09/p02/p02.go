package main

import "fmt"

const Pi float64 = 3.141592653589793

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * Pi * c.Radius
}

func main() {
	c := Circle{Radius: 1}

	fmt.Println(c.Area(), c.Perimeter())
}
