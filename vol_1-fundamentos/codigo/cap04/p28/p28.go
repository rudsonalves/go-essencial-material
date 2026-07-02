package main

import "fmt"

func countChar(str string) (rune, int) {
	count := map[rune]int{}

	for _, c := range str {
		count[c]++
	}

	r := ' '
	c := 0
	for k, v := range count {
		if v > c {
			r = k
			c = v
		}
	}

	return r, c
}

func main() {
	word := "banana"

	r, c := countChar(word)
	fmt.Printf("%c -> %d\n", r, c)
}
