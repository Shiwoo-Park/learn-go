package banking

import (
	"errors"
	"fmt"
)

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
- add pointer receiver: (a *Account)
- Shouldn't make copy of account, so *Account is important
*/
func (a *Account) Deposit(amount int) {
	fmt.Println("Deposit: ", amount)
	a.balance += amount
}

var errorNoMoney = errors.New("Can't Withdraw: Not enough balance")

// Withdraw from the account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errorNoMoney
	}

	fmt.Println("Withdraw: ", amount)
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

/*
Balance of account

- Since we don't use pointer receiver in this function,
- "a" in this method is a "copy" of the Account
*/
func (a Account) Balance() int {
	fmt.Println("Balance: ", a.balance)
	return a.balance
}

// Owner of the account
func (a Account) Owner() string {
	fmt.Println("Owner: ", a.owner)
	return a.owner
}

/*
String representation of Account struct

- similar with python __str__
*/
func (a Account) String() string {
	return fmt.Sprint(a.owner, "'s accunt. Has: ", a.balance)
}
