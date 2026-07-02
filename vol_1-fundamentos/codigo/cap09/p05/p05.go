package main

import (
	"fmt"
)

type Currency struct {
	Value int
}

func (c Currency) String() string {
	real := float64(c.Value) / 100.
	return fmt.Sprintf("R$ %.2f", real)
}

func (c Currency) Add(o Currency) Currency {
	return Currency{
		Value: c.Value + o.Value,
	}
}

func (c Currency) Sub(o Currency) Currency {
	return Currency{
		Value: c.Value - o.Value,
	}
}

type BankAccount struct {
	Balance Currency
}

func (b *BankAccount) Deposit(value Currency) {
	b.Balance = b.Balance.Add(value)
}

func (b *BankAccount) Withdraw(value Currency) bool {
	if b.Balance.Value < value.Value {
		return false
	}

	b.Balance = b.Balance.Sub(value)
	return true
}

func main() {
	account := BankAccount{
		Balance: Currency{},
	}

	fmt.Println("Saldo:", account.Balance)

	account.Deposit(Currency{25000})
	fmt.Println("Saldo:", account.Balance)

	account.Deposit(Currency{5000})
	fmt.Println("Saldo:", account.Balance)

	ok := account.Withdraw(Currency{20000})
	if !ok {
		fmt.Println("Saldo insuficiente...")
	} else {
		fmt.Println("Saldo:", account.Balance)
	}

	ok = account.Withdraw(Currency{20000})
	if !ok {
		fmt.Println("Saldo insuficiente...")
	} else {
		fmt.Println("Saldo:", account.Balance)
	}
}
