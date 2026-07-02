package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Atendente %d, processando cliente %d\n", id, job)
		delay := time.Duration(rand.IntN(1501)+500) * time.Millisecond
		time.Sleep(delay)
	}
}

func buildQueue(jobs chan<- int) {
	defer close(jobs)

	for i := range 20 {
		jobs <- i + 1
	}
}

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for i := range 4 {
		wg.Add(1)
		go worker(i+1, jobs, &wg)
	}

	go buildQueue(jobs)

	wg.Wait()
}
