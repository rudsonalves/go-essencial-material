package main

import (
	"fmt"
	"time"
)

func sendFast(ch chan<- string) {
	time.Sleep(2 * time.Second)
	ch <- "Mensagm fast"
}

func sendSlow(ch chan<- string) {
	time.Sleep(3 * time.Second)
	ch <- "Mensagm slow"
}

func main() {
	fash := make(chan string)
	slow := make(chan string)

	go sendSlow(slow)
	go sendFast(fash)

	select {
	case msg := <-fash:
		fmt.Println(msg)
	case msg := <-slow:
		fmt.Println(msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Tempo acabou!")
	}
}
