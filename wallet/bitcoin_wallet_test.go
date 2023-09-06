package wallet

import (
	"fmt"
	"testing"
)

func TestBitcoinWallet(t *testing.T) {
	t.Run("wallet should let us deposit Bitcoin and see the balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(.10))

		got := wallet.Balance()
		fmt.Printf("address of total in test is %v \n", &got)

		want := Bitcoin(0.10)

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})
	t.Run("wallet should let us Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(.10))
		wallet.Withdraw(Bitcoin(.05))

		got := wallet.Balance()
		want := Bitcoin(.05)

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})
}
