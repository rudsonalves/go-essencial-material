package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	CardNumber string
	Limit      float64
}

func (cc CreditCard) Pay(amount float64) error {
	if cc.CardNumber == "" {
		return errors.New("conta corrente sem cartão para pagamentos")
	}

	if cc.Limit == 0 {
		return errors.New("cartão sem limite aprovado")
	}

	if cc.Limit < amount {
		return errors.New("ultrapassou o limite de seu cartão")
	}

	fmt.Printf("Pagamento no cartão efetuado: R$ %.2f\n", amount)
	return nil
}

type Pix struct {
	Balance float64
}

func (p *Pix) Pay(amount float64) error {
	if p.Balance < amount {
		return errors.New("não possui saldo suficiente")
	}

	p.Balance -= amount
	fmt.Printf("Pix efetuado: R$ %.2f\n", amount)
	return nil
}

type Cash struct {
	Balance float64
}

func (c *Cash) Pay(amount float64) error {
	if c.Balance < amount {
		return errors.New("não possui saldo suficiente")
	}

	c.Balance -= amount
	fmt.Printf("Pagamento efetuado: R$ %.2f\n", amount)
	return nil
}

func Pay(p PaymentMethod, amount float64) error {
	return p.Pay(amount)
}

func main() {
	cc := &CreditCard{}
	p := &Pix{200}
	c := &Cash{200}

	if err := Pay(cc, 10); err != nil {
		fmt.Println(err)
	}

	if err := Pay(p, 10); err != nil {
		fmt.Println(err)
	}

	if err := Pay(c, 250); err != nil {
		fmt.Println(err)
	}

	cc.CardNumber = "123"
	if err := Pay(cc, 10); err != nil {
		fmt.Println(err)
	}

	cc.Limit = 100
	if err := Pay(cc, 120); err != nil {
		fmt.Println(err)
	}

	if err := Pay(cc, 10); err != nil {
		fmt.Println(err)
	}
}
