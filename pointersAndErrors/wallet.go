package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC\n", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("Not enough Bitcoins, Bro")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.Balance() < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
