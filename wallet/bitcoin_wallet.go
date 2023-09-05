package wallet

type Wallet struct {
	Total float64
}

func (w Wallet) Deposit(deposit float64) {
	w.Total += deposit
}

func (w Wallet) Balance() (balance float64) {
	return 0
}
