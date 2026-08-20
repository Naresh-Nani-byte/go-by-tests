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

	assertError := func (t testing.TB, err error)  {
		t.Helper()
		if err == nil{
			t.Errorf("wanted an error but didn't got")
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
		assertError(t,err)
		assertBalance(t, wallet, startingBalance)

	})
}