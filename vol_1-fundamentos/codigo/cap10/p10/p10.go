package main

import (
	"errors"
	"fmt"
)

type Notifies interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) error {

	fmt.Println("Email enviado...")
	return nil
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	if len(message) > 20 {
		return errors.New("mensagem não pode ter mais de 20 caracteres")
	}

	fmt.Println("SMS enviado...")
	return nil
}

func SendAll(notifies []Notifies, message string) {
	for _, n := range notifies {
		if err := n.Send(message); err != nil {
			fmt.Println("ERROR:", err)
		}
	}
}

func main() {
	e := &EmailNotifier{}
	s := &SMSNotifier{}

	message := "Salve os Golfinhos"
	SendAll([]Notifies{e, s}, message)

	fmt.Println()
	message = "A imaginação é mais importante que o conhecimento. O conhecimento é limitado, enquanto a imaginação abrange o mundo inteiro."
	SendAll([]Notifies{e, s}, message)
}
