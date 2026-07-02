package main

import "fmt"

type Statistics struct {
	Values []int
}

func (s Statistics) String() string {
	return fmt.Sprint(s.Values)
}

func (s Statistics) Len() int {
	return len(s.Values)
}

func (s Statistics) Min() int {
	min := s.Values[0]
	for _, v := range s.Values {
		if min < v {
			min = v
		}
	}

	return min
}

func (s Statistics) Max() int {
	max := s.Values[0]
	for _, v := range s.Values {
		if max > v {
			max = v
		}
	}

	return max
}

func (s Statistics) Average() float64 {
	sum := 0
	for _, v := range s.Values {
		sum += v
	}

	return float64(sum) / float64(s.Len())
}

func main() {
	s := Statistics{[]int{10, 20, 30, 40, 50, 60}}

	fmt.Println("Min:", s.Min())
	fmt.Println("Max:", s.Max())
	fmt.Println("Average:", s.Average())
}
