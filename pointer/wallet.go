package main

import "fmt"

type Bitcoin int

type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet ,,,
type Wallet struct {
	balance Bitcoin
}

// Deposit ,,,
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Print(w)
	// fmt.Print(*w)
	// fmt.Println("address of wallet in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance ,,,
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
