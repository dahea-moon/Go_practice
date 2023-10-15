package main

import (
	"fmt"

	"github.com/dahea-moon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("rachel")
	account.Deposit(10)
	fmt.Println(account.Balance())
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance(), account.Owner())
	account.ChangeOwner("new rachel")
	fmt.Println(account)
}
