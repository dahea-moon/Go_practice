package main

import (
	"fmt"

	"github.com/dahea-moon/learngo/accounts"
)

func main() {
	account := accounts.NewAccount("rachel")
	fmt.Println(account)
}
