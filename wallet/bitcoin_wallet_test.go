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

	assertError := func(t testing.TB, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error, but expected one")
		}

		if got.Error() != want.Error() {
			t.Errorf("got %q, but wanted %q", got, want)
		}
	}

	assertNoError := func(t testing.TB, err error) {
		t.Helper()
		if err != nil {
			t.Errorf("got error, but didn't expect one")
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
		err := wallet.Withdraw(Bitcoin(5))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})
	t.Run("withdraw refused due to insufficient funds", func(t *testing.T) {
		wallet := Wallet{total: Bitcoin(5)}
		got := wallet.Withdraw(Bitcoin(10))
		want := ErrorInsufficientFunds

		assertError(t, got, want)

	})
}
