package geometry

type Shape interface {
	Area() float64
	Perimeter() float64
}

func Area(s Shape) float64 {
	return s.Area()
}

func Perimeter(s Shape) float64 {
	return s.Perimeter()
}
