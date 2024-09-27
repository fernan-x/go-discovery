package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("wallet deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("wallet withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(5)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw insuficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(10)
		wallet := Wallet{initialBalance}
		err := wallet.Withdraw(20)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, initialBalance)
	})
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("got %q want nil", err)
	}
}

func assertBalance(t *testing.T, wallet Wallet, expected Bitcoin) {
	t.Helper()
	result := wallet.Balance()
	if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}
}

func assertError(t *testing.T, err error, expected error) {
	t.Helper()
	if err == nil {
		t.Errorf("wanted an error but didn't get one")
	}

	if err != expected {
		t.Errorf("got %q want %q", err.Error(), expected)
	}
}