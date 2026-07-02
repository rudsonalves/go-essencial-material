package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.Width + r.Height)
}

func main() {
	r := Rectangle{10, 5}

	fmt.Println("Área:", r.Area())
	fmt.Println("Perímetro:", r.Perimeter())
}
