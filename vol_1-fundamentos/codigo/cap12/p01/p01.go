package main

import (
	"fmt"

	"p01/mathutil"
)

func main() {
	i, j := 2, 6

	max := mathutil.Max(i, j)
	min := mathutil.Min(i, j)

	fmt.Println("Máximo:", max)
	fmt.Println("Mínimo:", min)
}
