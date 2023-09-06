package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	total Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.total += deposit
}

func (w *Wallet) Withdraw(withdraw Bitcoin) {
	w.total -= withdraw
}

func (w *Wallet) Balance() Bitcoin {
	return 0
}

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
