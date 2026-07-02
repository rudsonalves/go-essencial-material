package solver

import (
	"errors"
	"math"
)

func Solver(a, b, c float64) (float64, float64, error) {
	d := delta(a, b, c)
	if d < 0 {
		return 0, 0, errors.New("não possui raizes reais")
	}

	a2 := 2 * a

	sq := math.Sqrt(d) / a2

	return -b/a2 + sq, -b/a2 - sq, nil
}

func delta(a, b, c float64) float64 {
	return b*b - 4*a*c
}
