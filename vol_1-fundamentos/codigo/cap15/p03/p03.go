package main

import (
	"fmt"
	"sync"
)

func counter(workId int, c *int, wg *sync.WaitGroup, mu *sync.Mutex) {
	defer wg.Done()

	for range 100 {
		mu.Lock()
		(*c)++
		mu.Unlock()
		fmt.Printf("%d - count: %d\n", workId, *c)
	}
}

func main() {
	var (
		c  int
		wg sync.WaitGroup
		mu sync.Mutex
	)

	for i := range 10 {
		wg.Add(1)

		go counter(i+1, &c, &wg, &mu)
	}

	wg.Wait()
	fmt.Println(c)
}
