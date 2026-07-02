package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func delay() time.Duration {
	return time.Duration(rand.IntN(5)+1) * time.Second
}

func routine(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("goroutine %d iniciada\n", id)
	time.Sleep(delay())
	fmt.Printf("goroutine %d finalizada\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := range 3 {
		wg.Add(1)
		go routine(i+1, &wg)
	}

	wg.Wait()
}
