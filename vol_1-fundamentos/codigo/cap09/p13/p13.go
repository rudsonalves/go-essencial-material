package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func (p Point) Distance() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func main() {
	p := Point{10, 20}

	fmt.Println(p)
	fmt.Println(p.Distance())

	p.Move(10, 20)
	fmt.Println(p)
	fmt.Println(p.Distance())
}
