package wallet

import "fmt"

type Wallet struct {
	total float64
}

func (w *Wallet) Deposit(deposit float64) {
	w.total += deposit
	fmt.Printf("address of total in main file is %v \n", &w.total)
}

func (w *Wallet) Balance() float64 {
	fmt.Printf("address of total in Ballance in main file is %v \n", &w.total)
	return w.total
}
