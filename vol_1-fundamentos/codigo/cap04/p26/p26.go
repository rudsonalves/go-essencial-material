package main

import "fmt"

func main() {
	frase := "banana"
	count := map[rune]int{}

	for _, c := range frase {
		// _, ok := count[c]
		// if ok {
		count[c]++
		// } else {
		// 	count[c] = 1
		// }
	}

	for key, value := range count {
		fmt.Printf("%c -> %d\n", key, value)
	}
}
