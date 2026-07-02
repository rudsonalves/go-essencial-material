package main

import (
	"fmt"
)

func produce(ch chan<- int) {
	defer close(ch)

	for i := range 5 {
		ch <- i + 1
	}
}

func main() {
	ch := make(chan int)

	go produce(ch)

	for i := range ch {
		fmt.Println(i)
	}
}
