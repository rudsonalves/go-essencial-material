package calculator

import "testing"

func TestSum(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"positive numbers", 2, 3, 5},
		{"positive numbers", 10, 5, 15},
		{"different signs", -2, 2, 0},
		{"negative numbers", -5, -8, -13},
		{"ero values", 0, 0, 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			got := Sum(test.a, test.b)

			if got != test.want {
				t.Errorf(
					"Sum(%d, %d) = %d; want %d",
					test.a,
					test.b,
					got,
					test.want,
				)
			}
		})
	}
}

func BenchmarkSum(b *testing.B) {
	a := 2
	c := 3

	b.ResetTimer()

	for range b.N {
		Sum(a, c)
	}
}
