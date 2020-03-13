package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		// fmt.Println("type", reflect.TypeOf(wallet))

		wallet.Deposit(Bitcoin(10))

		// fmt.Println("address of wallet in test is", &wallet.balance)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	t.Run("test getting valur", func(t *testing.T) {
		wallet := Wallet{}

		walletPointer := &wallet

		fmt.Println(wallet.Balance())

		fmt.Println(walletPointer.Balance())

	})
}
