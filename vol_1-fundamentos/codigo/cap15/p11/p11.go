package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Philosopher struct {
	Name              string
	RemainingPortions int // food portions
	Portions          int
}

func NewPhilosopher(name string) Philosopher {
	portions := rand.IntN(4) + 7

	return Philosopher{
		Name:              name,
		RemainingPortions: portions,
		Portions:          portions,
	}
}

func (p *Philosopher) Eat() bool {
	if p.RemainingPortions == 0 {
		return false
	}

	p.RemainingPortions--
	time.Sleep(randomDelay())

	return true
}

func (p Philosopher) Finished() bool {
	return p.RemainingPortions == 0
}

func randomDelay() time.Duration {
	return time.Duration(rand.IntN(1501)+500) * time.Millisecond
}

func (p Philosopher) String() string {
	return fmt.Sprintf("%s (%d/%d porções)", p.Name, p.RemainingPortions, p.Portions)
}

func diner(p *Philosopher, fork chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	for !p.Finished() {
		fmt.Printf("%s está pensando...\n", p.Name)
		time.Sleep(randomDelay())

		fmt.Printf("%s aguardando o garfo...\n", p.Name)
		<-fork

		fmt.Printf("%s pegou o garfo\n", p.Name)
		p.Eat()
		fmt.Printf("%s comeu uma porção (%d/%d porções)\n", p.Name, p.RemainingPortions, p.Portions)

		fork <- struct{}{}
		fmt.Printf("%s devolveu o garfo\n", p.Name)
	}

	fmt.Printf("%s terminou sua canja\n", p.Name)
}

func main() {
	var wg sync.WaitGroup
	fork := make(chan struct{}, 1)

	s := NewPhilosopher("Sócrates")
	p := NewPhilosopher("Platão")
	a := NewPhilosopher("Aristóteles")

	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(a)
	fmt.Println()

	wg.Add(3)
	go diner(&s, fork, &wg)
	go diner(&p, fork, &wg)
	go diner(&a, fork, &wg)

	fork <- struct{}{}
	wg.Wait()
}
