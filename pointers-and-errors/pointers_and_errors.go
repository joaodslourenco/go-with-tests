package pointersanderrors

import (
	"errors"
	"fmt"
)

var ErrInsuficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.Balance() {
		return ErrInsuficientFunds
	}

	w.balance -= value
	return nil
}
