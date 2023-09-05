package wallet

import (
	"fmt"
	"testing"
)

func TestBitcoinWallet(t *testing.T) {
	t.Run("wallet should let us deposit Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(0.10)

		got := wallet.Balance()
		fmt.Printf("address of total in test is %v \n", &got)

		want := 0.10

		if got != want {
			t.Errorf("got %g, but wanted %g", got, want)
		}
	})
}
