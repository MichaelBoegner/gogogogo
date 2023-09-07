package wallet

import (
	"fmt"
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
	t.Run("withdraw refused due to insufficient funds", func(t *testing.T) {
		wallet := Wallet{total: Bitcoin(5)}
		got := wallet.Withdraw(Bitcoin(10))
		want := fmt.Sprintf("Insufficient funds. There are %s left in the account.", wallet.total)

		if got != want {
			t.Errorf("got %s, but wanted %s", got, want)
		}
	})
}
