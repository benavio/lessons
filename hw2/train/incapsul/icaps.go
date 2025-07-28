package main

// ИНКАПСУЛЯЦИЯ

import "fmt"

type Wallet struct {
	balance float64
}

func (w *Wallet) Deposit(amount float64) {
	w.balance = w.balance + amount
}
func (w *Wallet) Withdraw(amount float64) {
	w.balance = w.balance - amount
}
func (w *Wallet) Balance() {
	fmt.Printf("i've got %v balance", w.balance)
	fmt.Print(" ")
}

func main() {
	money := Wallet{100}

	money.Deposit(13.2)
	money.Balance()
	money.Withdraw(92.1)
	money.Balance()

}
