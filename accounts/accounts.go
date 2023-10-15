package accounts

import (
	"errors"
	"fmt"
)

// Account private struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccoutn create Account
func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account method
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account method
func (a account) Balance() int {
	return a.balance
}

// withdraw x from your account method
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// get owner of the account
func (a *account) Owner() string {
	return a.owner
}

func (a *account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \n Has: ", a.Balance())
}
