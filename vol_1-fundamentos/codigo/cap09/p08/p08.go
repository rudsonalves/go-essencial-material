package main

import "fmt"

type Stack struct {
	Values []int
}

func (s Stack) String() string {
	return fmt.Sprint(s.Values)
}

func (s *Stack) Push(value int) {
	s.Values = append(s.Values, value)
}

func (s *Stack) Pop() (int, bool) {
	index := len(s.Values) - 1
	if index < 0 {
		return 0, false
	}

	value := s.Values[index]
	s.Values = s.Values[:index][:]

	return value, true
}

func main() {
	a := Stack{}

	fmt.Println("Carregando a pilha...")
	for i := range 10 {
		a.Push(2 * i)
		fmt.Println(a)
	}

	fmt.Println("\nDesarrgando a pilha...")

	for {
		value, ok := a.Pop()
		if ok {
			fmt.Println(value)
		} else {
			fmt.Println("Pilha vazia!")
			break
		}
	}
}
