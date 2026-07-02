package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	m := rand.IntN(12)

	month := []string{
		"Jan",
		"Fev",
		"Mar",
		"Abr",
		"Mai",
		"Jun",
		"Jul",
		"Ago",
		"Set",
		"Out",
		"Nov",
		"Dez",
	}

	fmt.Printf("%d - %s\n", m+1, month[m])
}
