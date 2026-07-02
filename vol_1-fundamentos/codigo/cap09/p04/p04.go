package main

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() int {
	c.Value++
	return c.Value
}

func (c *Counter) Decrement() int {
	c.Value--
	return c.Value
}

func main() {
	c := Counter{5}

	fmt.Println(c.Increment())
	fmt.Println(c.Increment())
	fmt.Println(c.Increment())
	fmt.Println(c.Increment())
	fmt.Println(c.Decrement())
	fmt.Println(c.Decrement())
}
