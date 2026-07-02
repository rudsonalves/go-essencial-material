package main

import "fmt"

func findInB(x int, b []int) bool {
	for _, v := range b {
		if v == x {
			return true
		}
	}

	return false
}

func isEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for _, vA := range a {
		if !findInB(vA, b) {
			return false
		}
	}

	return true
}

func main() {
	a := []int{3, 7, 2, 5}
	b := []int{5, 2, 7, 3}

	const (
		msgFailure string = "Os slices são diferentes"
		msgSuccess string = "Os slices são iguais"
	)

	var msg string = ""
	if isEquals(a, b) {
		msg = msgSuccess
	} else {
		msg = msgFailure
	}

	fmt.Println(msg)
}
