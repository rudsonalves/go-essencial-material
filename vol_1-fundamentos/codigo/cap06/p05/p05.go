package main

import "fmt"

func concat(words ...string) string {
	str := ""
	for _, w := range words {
		space := ""
		if len(str) > 0 {
			space = " "
		}
		str += space + w
	}

	return str
}

func main() {
	s := concat("Go", "é", "simples")

	fmt.Println(s)
}
