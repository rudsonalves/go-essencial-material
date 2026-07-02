package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

var _ Shape = Circle{}
var _ Shape = Rectangle{}

func main() {
	PrintArea(Circle{Radius: 3})
	PrintArea(Rectangle{Width: 4, Height: 5})
}
