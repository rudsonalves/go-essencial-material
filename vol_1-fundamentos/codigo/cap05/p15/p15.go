package main

import "fmt"

func hasKey(set map[int]int, key int) bool {
	for k, _ := range set {
		if k == key {
			return true
		}
	}

	return false
}

func main() {
	numbers := []int{3, 3, 3, 7, 7, 7, 7, 7, 2, 5, 5, 5, 5, 1}

	countSeq := map[int]int{}

	for i, v := range numbers {
		if hasKey(countSeq, v) {
			continue
		}
		count := 1
		for j := i + 1; j < len(numbers); j++ {
			if v == numbers[j] {
				count++
				continue
			}

			countSeq[v] = count
			break
		}
	}

	max := 0
	key := 0
	for k, v := range countSeq {
		if v > max {
			max = v
			key = k
		}
	}

	fmt.Printf("%d -> %d\n", key, countSeq[key])
}
