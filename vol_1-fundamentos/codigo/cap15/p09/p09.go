package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Atendente %d, processando cliente %d\n", id, job)

		delay := time.Duration(rand.IntN(1501)+500) * time.Millisecond
		time.Sleep(delay)

		results <- fmt.Sprintf("Cliente %d -> %d", job, 2*job)
	}
}

func buildQueue(jobs chan<- int) {
	defer close(jobs) // jobs devem ser fechados após todas as tarefas forem criadas

	for i := range 20 {
		jobs <- i + 1
	}
}

func closeAll(results chan string, wg *sync.WaitGroup) {
	wg.Wait()
	close(results) // aguarda que todos os resultados tenham sido calculados para fechar results
}

func main() {
	jobs := make(chan int)
	results := make(chan string)

	var workersWG sync.WaitGroup

	for i := range 4 {
		workersWG.Add(1)
		go worker(i+1, jobs, results, &workersWG)
	}

	go buildQueue(jobs)

	go closeAll(results, &workersWG)

	for result := range results {
		fmt.Println(result)
	}

}
