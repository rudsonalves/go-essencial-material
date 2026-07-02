package main

import (
	"fmt"
	"p02/geometry"
)

func main() {
	l := 3.
	c := geometry.Circle{Radius: l}
	s := geometry.Square{Side: 2 * l}
	r := geometry.Rectangle{Height: 1.5 * l, Width: 8.}

	fmt.Printf(
		"Círculo de raio %.2f:\n Área: %.2f\n Perímetro: %.2f\n\n",
		l,
		geometry.Area(c),
		geometry.Perimeter(c),
	)

	fmt.Printf(
		"Quadrado de aresta %.2f:\n Área: %.2f\n Perímetro: %.2f\n\n",
		2*l,
		geometry.Area(s),
		geometry.Perimeter(s),
	)

	fmt.Printf(
		"Retângulo de %.2f x %.2f:\n Área: %.2f\n Perímetro: %.2f\n\n",
		1.5*l, 8.,
		geometry.Area(r),
		geometry.Perimeter(r),
	)
}
