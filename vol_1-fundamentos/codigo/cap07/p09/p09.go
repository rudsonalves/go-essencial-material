package main

import "fmt"

func divide(a, b int, q, r *int, ok *bool) {
	if b == 0 {
		*ok = false
		return
	}

	*q = a / b
	*r = a % b
	*ok = true
}

func main() {
	a, b := 10, 3
	var q, r int
	var ok bool = false

	divide(a, b, &q, &r, &ok)
	if !ok {
		fmt.Printf("%d/%d não permitida!\n", a, b)
	} else {
		fmt.Printf("%d/%d = %d con resto %d\n", a, b, q, r)
	}

	a, b = 5, 0
	divide(a, b, &q, &r, &ok)
	if !ok {
		fmt.Printf("%d/%d não permitida!\n", a, b)
	} else {
		fmt.Printf("%d/%d = %d con resto %d\n", a, b, q, r)
	}

}
