package main

import "fmt"

func split(str string) []string {
	words := []string{}
	start := 0

	for i, c := range str {
		if c == 32 {
			if start == i {
				continue
			}

			words = append(words, str[start:i])
			start = i + 1
		}
	}

	words = append(words, str[start:])
	return words
}

func main() {
	frase := "go é simples e go é rápido"

	words := split(frase)
	count := map[string]int{}

	for _, word := range words {
		// if _, ok := count[word]; ok {
		count[word]++
		// } else {
		// 	count[word] = 1
		// }
	}

	for key, value := range count {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
