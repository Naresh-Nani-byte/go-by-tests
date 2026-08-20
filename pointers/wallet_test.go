package pointers

import "testing"

func TestWallet(t *testing.T){

	t.Run("Deposit the Bitcoin ", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		got := wallet.Balance()
		want := 10

		if got != want {
			t.Errorf("got %d but want %d ", got, want)
		}
	})	
}