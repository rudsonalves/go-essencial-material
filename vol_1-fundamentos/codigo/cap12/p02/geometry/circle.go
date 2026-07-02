package geometry

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return PI * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * PI * c.Radius
}
