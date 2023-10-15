package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"kimchi"}
	rachel := person{name: "rachel", age: 18, favFood: favFood}
	fmt.Println(rachel.age)
}
