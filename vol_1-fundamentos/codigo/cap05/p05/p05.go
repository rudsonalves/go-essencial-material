package main

import "fmt"

func main() {
	numbers := []int{7, 3, 5, 7, 2, 5}

	mapa := map[int]int{}

	for index, value := range numbers {
		_, ok := mapa[value]
		if !ok {
			mapa[value] = index
		}
	}

	for key, value := range mapa {
		fmt.Printf("%d -> %d\n", key, value)
	}
}
