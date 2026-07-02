package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func delay() time.Duration {
	return time.Duration(rand.IntN(5)+1) * time.Second
}

func routine(id int) {
	fmt.Printf("goroutine %d iniciada\n", id)
	time.Sleep(delay())
	fmt.Printf("goroutine %d finalizada\n", id)
}

func main() {
	for i := range 3 {
		go routine(i + 1)
	}

	time.Sleep(5 * time.Second)
}
