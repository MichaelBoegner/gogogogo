package wallet

import (
	"fmt"
	"testing"
)

func TestBitcoinWallet(t *testing.T) {
	t.Run("wallet should let us deposit Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(.10))

		got := wallet.Balance()
		fmt.Printf("address of total in test is %v \n", &got)

		want := Bitcoin(0.10)

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})
}
