package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	n := rand.IntN(7) + 1
	weekDay := []string{"Domingo", "Segunda", "Terça", "Quarta", "Quinta", "Sexta", "Sábado"}

	fmt.Printf("%d - %s\n", n, weekDay[n-1])
}
