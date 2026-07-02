package main

import "fmt"

type Logger interface {
	Log(message string)
}

type Terminal struct{}

func (t Terminal) Log(message string) {
	fmt.Println(message)
}

type Memory struct {
	Data string
}

func (m *Memory) Log(message string) {
	m.Data = message
}

func Log(l Logger, message string) {
	l.Log(message)
}

func main() {
	t := Terminal{}
	m := &Memory{}

	Log(t, "Imprime no terminal...")
	Log(m, "Armazenado na memória...")

	fmt.Println(">>>", m.Data)
}
