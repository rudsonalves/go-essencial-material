package mathutil

import (
	"testing"
)

func TestMax(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positivos 1", 2, 3, 3},
		{"positivos 2", 3, 2, 3},
		{"negativos 1", -2, -3, -2},
		{"negativos 2", -3, -2, -2},
		{"mistos 1", -2, 3, 3},
		{"mistos 2", -3, 2, 2},
		{"iguais", 2, 2, 2},
		{"zeros", 0, 3, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Max(test.a, test.b)

			if got != test.want {
				t.Errorf("Max(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positivos 1", 2, 3, 2},
		{"positivos 2", 3, 2, 2},
		{"negativos 1", -2, -3, -3},
		{"negativos 2", -3, -2, -3},
		{"mistos 1", -2, 3, -2},
		{"mistos 2", -3, 2, -3},
		{"iguais", 2, 2, 2},
		{"zeros", 0, 3, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Min(test.a, test.b)

			if got != test.want {
				t.Errorf("Min(%d, %d) = %d; want %d", test.a, test.b, got, test.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"positivo 1", 2, 2},
		{"positivo 2", 2, 2},
		{"negativo 1", -2, 2},
		{"negativo 2", -3, 3},
		{"zero", 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := Abs(test.n)

			if got != test.want {
				t.Errorf("Abs(%d) = %d; want %d", test.n, got, test.want)
			}
		})
	}
}
