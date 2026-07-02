package main

import "fmt"

type Timer struct {
	Seconds int
}

func (t Timer) String() string {
	return fmt.Sprintf("%d sec", t.Seconds)
}

func (t *Timer) Add(seconds int) {
	t.Seconds += seconds
}

func (t *Timer) Reset() {
	t.Seconds = 0
}

func main() {
	timer := Timer{}

	timer.Add(30)
	timer.Add(45)

	fmt.Println(timer.String())

	timer.Reset()

	fmt.Println(timer.String())
}
