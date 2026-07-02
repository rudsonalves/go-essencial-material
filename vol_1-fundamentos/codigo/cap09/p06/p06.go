package main

import "fmt"

type Student struct {
	Score [3]float64
}

func (s Student) Average() float64 {
	sum := 0.0
	for _, s := range s.Score {
		sum += s
	}

	return sum / 3.
}

func (s Student) Approved() bool {
	if s.Average() >= 6 {
		return true
	}

	return false
}

func main() {
	myClass := map[string]Student{
		"Maria":    Student{Score: [3]float64{6., 5.5, 8.0}},
		"Sérgio":   Student{Score: [3]float64{2, 5, 5}},
		"Fernanda": Student{Score: [3]float64{8, 9.5, 10}},
	}

	for name, st := range myClass {
		status := "reprovado"
		if st.Approved() {
			status = "aprovado"
		}
		fmt.Printf("%10s: %2.1f (%s)\n", name, st.Average(), status)
	}
}
