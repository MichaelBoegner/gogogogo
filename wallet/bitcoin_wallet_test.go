package wallet

import "testing"

func TestBitcoinWallet(t *testing.T) {
	t.Run("wallet should let us deposit Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(0.10)

		got := wallet.Balance()
		want := 0.10

		if got != want {
			t.Errorf("got %g, but wanted %g", got, want)
		}
	})
}
