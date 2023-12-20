package structs

import (
	"errors"
	"fmt"
	"hello/testhelper"
	"testing"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var InsufficentFundsError = "insufficient funds"

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if (w.balance - amount) < 0 {
		return errors.New(InsufficentFundsError)
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) GetBalance() Bitcoin {
	return w.balance
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		r := Wallet{}
		r.Deposit(Bitcoin(10))
		got := r.GetBalance()
		wanted := Bitcoin(10)
		if got != wanted {
			t.Errorf("%#v expected %v but got %v", got, got, wanted)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		r := Wallet{20}
		err := r.Withdraw(Bitcoin(14))
		testhelper.AssertNoError(t, err)
		wanted := Bitcoin(6)
		got := r.balance
		if got != wanted {
			t.Errorf("%#v expected %v but got %v", got, got, wanted)
		}
	})

	t.Run("overdraw", func(t *testing.T) {
		r := Wallet{20}
		err := r.Withdraw(Bitcoin(200))
		testhelper.AssertError(t, err, InsufficentFundsError)
	})

}
