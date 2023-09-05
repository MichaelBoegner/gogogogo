package wallet

import "fmt"

type Bitcoin float64

type Wallet struct {
	total Bitcoin
}

func (w *Wallet) Deposit(deposit Bitcoin) {
	w.total += deposit
	fmt.Printf("address of total in main file is %v \n", &w.total)
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of total in Ballance in main file is %v \n", &w.total)
	return w.total
}
