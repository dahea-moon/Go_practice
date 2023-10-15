package main

import (
	"fmt"

	"github.com/dahea-moon/learngo/mydict"
	// "github.com/dahea-moon/learngo/accounts"
)

func main() {
	dictionary := mydict.Dictionary{"first": "First word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
}
