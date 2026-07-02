package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Philosopher struct {
	Name         string
	Plate        int // food portions
	InitialPlate int
}

func NewPhilosopher(name string) Philosopher {
	portions := rand.IntN(4) + 7

	return Philosopher{
		Name:         name,
		Plate:        portions,
		InitialPlate: portions,
	}
}

func (p *Philosopher) Eat() bool {
	if p.Plate == 0 {
		return false
	}

	p.Plate--
	time.Sleep(randomDelay())

	return true
}

func (p Philosopher) Finished() bool {
	return p.Plate == 0
}

func randomDelay() time.Duration {
	return time.Duration(rand.IntN(1501)+500) * time.Millisecond
}

func (p Philosopher) String() string {
	return fmt.Sprintf("%s (%d/%d porções)", p.Name, p.Plate, p.InitialPlate)
}

func diner(p *Philosopher, fork *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for !p.Finished() {
		fmt.Printf("%s está pensando...\n", p.Name)
		time.Sleep(randomDelay())

		fmt.Printf("%s aguardando o garfo...\n", p.Name)
		fork.Lock()

		fmt.Printf("%s pegou o garfo\n", p.Name)
		p.Eat()
		fmt.Printf("%s comeu uma porção (%d/%d porções)\n", p.Name, p.Plate, p.InitialPlate)

		fork.Unlock()
		fmt.Printf("%s devolveu o garfo\n", p.Name)
	}

	fmt.Printf("%s terminou sua canja\n", p.Name)
}

func main() {
	var (
		fork sync.Mutex
		wg   sync.WaitGroup
	)

	s := NewPhilosopher("Sócrates")
	p := NewPhilosopher("Platão")
	a := NewPhilosopher("Aristóteles")

	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(a)
	fmt.Println()

	wg.Add(3)
	go diner(&s, &fork, &wg)
	go diner(&p, &fork, &wg)
	go diner(&a, &fork, &wg)

	wg.Wait()
}
