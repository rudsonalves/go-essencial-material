package main

import (
	"fmt"
	"strings"
)

const desloc = 'a' - 'A'

type Text struct {
	Value string
}

func (t Text) String() string {
	return t.Value
}

func (t Text) Upper() string {
	result := ""

	for _, r := range t.Value {
		if r >= 'a' && r <= 'z' {
			r -= desloc
		}

		// Para fins didáticos utilizamos concatenação de strings.
		// Em aplicações reais, strings.Builder costuma ser mais eficiente.
		result += string(r)
	}

	return result
}

func (t Text) Lower() string {
	// strings.Builder é frequentemente utilizado quando são
	// necessárias muitas concatenações de strings.
	// Seu zero value já está pronto para uso.
	var result strings.Builder

	for _, r := range t.Value {
		if r >= 'A' && r <= 'Z' {
			r += desloc
		}

		result.WriteRune(r)
	}

	return result.String()
}

func (t Text) Find(v string) int {
	lenV := len(v)
	lenT := len(t.Value)

	if lenV == 0 {
		return 0
	}

	if lenV > lenT {
		return -1
	}

	for index := 0; index <= lenT-lenV; index++ {
		if t.Value[index:index+lenV] == v {
			return index
		}
	}

	return -1
}

func (t *Text) Replace(old, new string) bool {
	index := t.Find(old)
	if index < 0 {
		return false
	}

	t.Value = t.Value[:index] + new + t.Value[index+len(old):]
	return true
}

func main() {
	a := Text{"Alberto Santos"}

	fmt.Println(a)
	upper := a.Upper()
	fmt.Println(upper)

	lower := a.Lower()
	fmt.Println(lower)

	fmt.Println(a.Find("Alb"))
	fmt.Println(a.Find("San"))
	fmt.Println(a.Find("Santos"))

	a.Replace("S", "s")
	fmt.Println(a)

	ok := a.Replace("sant", "Pinheir")
	if !ok {
		fmt.Println("Não pode substituir!")
	} else {
		fmt.Println(a)
	}
}
