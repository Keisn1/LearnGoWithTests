package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Test Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("Test WithDraw happy path", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("Test WithDraw Not enough money", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		wantErr := ErrInsufficientFunds
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, wantErr)
	})
}

func assertError(t testing.TB, gotErr error, wantErr error) {
	t.Helper()
	if gotErr == nil {
		t.Fatal("Didn't get an error, but want an error to be thrown")
	}
	if gotErr != wantErr {
		t.Errorf("Got \"%s\"; want \"%s\"", gotErr, wantErr)
	}
}

func assertNoError(t testing.TB, gotErr error) {
	t.Helper()
	if gotErr != nil {
		t.Fatal("Got an Error, but didn't want one")
	}
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper() // for the line number to be shown up correctly

	got := wallet.Balance()
	if got != want {
		t.Errorf("wallet.Balance() = \"%s\"; want \"%s\"", got, want)
	}
}
