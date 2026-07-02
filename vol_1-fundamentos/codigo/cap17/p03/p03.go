package main

import "fmt"

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct{}

func (e EmailNotifier) Notify(message string) error {
	fmt.Printf("Notificação por Email: %s\n", message)

	return nil
}

type SMSNotifier struct{}

func (s SMSNotifier) Notify(message string) error {
	fmt.Printf("Notificação por SMS: %s\n", message)

	return nil
}

var _ Notifier = EmailNotifier{}

var _ Notifier = SMSNotifier{}

func main() {
	email := EmailNotifier{}
	sms := SMSNotifier{}

	email.Notify("Pedido confirmado")
	sms.Notify("Pedido enviado")
}
