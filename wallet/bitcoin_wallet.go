package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	total Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.total += deposit
}

func (w *Wallet) Withdraw(withdraw Bitcoin) (err string) {
	if withdraw > w.total {
		err := fmt.Sprintf("Insufficient funds. There are %s left in the account.", w.total)
		return err
	}
	w.total -= withdraw
	return err
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
