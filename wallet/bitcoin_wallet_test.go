package wallet

import (
	"testing"
)

func TestBitcoinWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	}

	t.Run("wallet should let us deposit Bitcoin and see the balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("wallet should let us Withdraw Bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
	})
}
