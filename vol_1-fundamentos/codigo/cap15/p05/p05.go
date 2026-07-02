package main

import "fmt"

func main() {
	ch := make(chan string, 3)

	for i := range 3 {
		ch <- fmt.Sprintf("Mensagem %d", i+1)
	}

	for range 3 {
		fmt.Println(<-ch)
	}
}
