package main

import (
	"fmt"

	"github.com/dahea-moon/learngo/mydict"
	// "github.com/dahea-moon/learngo/accounts"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseWord := "hello"
	dictionary.Add(baseWord, "First")
	err := dictionary.Update(baseWord, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseWord)
	fmt.Println(word)

}
