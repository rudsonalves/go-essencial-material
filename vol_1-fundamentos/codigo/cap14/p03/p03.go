package main

import (
	"fmt"
	"strings"
)

func clearSpaces(value string) string {
	value = strings.TrimSpace(value)

	for {
		newValue := strings.ReplaceAll(value, "  ", " ")
		if newValue == value {
			return value
		}
		value = newValue
	}
}

func capitalizeWord(word string) string {
	if word == "" {
		return word
	}

	word = strings.ToLower(word)

	first := strings.ToUpper(word[0:1])
	rest := word[1:]

	return first + rest
}

func capitalize(phrase string) string {
	clean := clearSpaces(phrase)

	words := strings.Split(clean, " ")
	for i, word := range words {
		words[i] = capitalizeWord(word)
	}

	return strings.Join(words, " ")
}

// A função strings.Fields remove automaticamente os espaços
// excedentes entre as palavras, tornando a implementação mais simples,
// sem a necessidade da função clearSpaces.
func capitalize2(phrase string) string {
	words := strings.Fields(phrase)

	for i, word := range words {
		words[i] = capitalizeWord(word)
	}

	return strings.Join(words, " ")
}

func main() {
	input := "   rObErTo           sAnToS   "

	cap := capitalize(input)
	fmt.Printf(">%s<\n", cap)

	cap2 := capitalize2(input)
	fmt.Printf(">%s<\n", cap2)
}
