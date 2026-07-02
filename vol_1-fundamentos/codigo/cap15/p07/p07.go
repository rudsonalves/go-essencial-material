package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"time"
)

func countDown(start int, ctx context.Context) {
	for i := range start {
		fmt.Println(start - i)

		select {
		case <-time.After(1 * time.Second):
			//
		case <-ctx.Done():
			fmt.Println("Operação cancelada:", ctx.Err())
			return
		}
	}
}

func main() {
	timeout := time.Duration(rand.IntN(8)+1) * time.Second

	fmt.Println("Timeout:", timeout)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		timeout,
	)
	defer cancel()

	countDown(10, ctx)
}
