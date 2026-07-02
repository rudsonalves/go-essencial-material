package main

import "fmt"

func countChar(p string, r rune) int {
	count := 0

	for _, c := range p {
		if c == r {
			count++
		}
	}

	return count
}

func main() {
	phrase := "bananas"

	charMap := map[rune]int{}

	for _, r := range phrase {
		if countChar(phrase, r) == 1 {
			charMap[r] = 1
		}
	}

	for k, v := range charMap {
		fmt.Printf("%c -> %d\n", k, v)
	}
}
