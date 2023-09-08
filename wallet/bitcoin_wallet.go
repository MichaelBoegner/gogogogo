package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	total Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.total += deposit
}

var ErrorInsufficientFunds = errors.New("cannot withdraw due to insufficient funds")

func (w *Wallet) Withdraw(withdraw Bitcoin) error {
	if withdraw > w.total {
		return ErrorInsufficientFunds
	}

	w.total -= withdraw
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.total
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
