package banking

import "fmt"

type Account struct {
	// public
	CustomName string

	// private
	owner   string
	balance int // default: 0
}

// NewAccount creates account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

/*
Deposit x amount money to account

- Account method
- add receiver: (a Account)
*/
func (a Account) Deposit(amount int) {
	fmt.Println("Deposit: ", amount)
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}
