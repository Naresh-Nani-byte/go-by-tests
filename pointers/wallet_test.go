package pointers

import (
	"testing"
)

func TestWallet(t *testing.T){

	assertBalance := func (t testing.TB, wallet Wallet, want Bitcoin)  {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d but want %d", got, want)
		}
	}

	assertError := func (t testing.TB, got error, want string)  {
		t.Helper()
		if got == nil{
			t.Fatal("wanted an error but didn't got")
		}
		if got.Error() != want {
			t.Errorf("got %q but want %q", got, want)
		}
		
	}
	t.Run("Deposit the Bitcoin ", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		assertBalance(t, wallet, 20)
	})	

	t.Run("Bitcoing String ", func(t *testing.T) {
		btc := Bitcoin(10)
		got := btc.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s but want %s",got, want)
		}
	})

	t.Run("Wothdraw the bitcoing", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw more than the capacity", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertError(t, err, "cannot withdraw, insufficient funds")
		assertBalance(t, wallet, startingBalance)

	})
}