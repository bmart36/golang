package main

import "fmt"

// We creatr VALUES fo a certain TYPE that are stored in VARIABLES
// and those VARIABLES have IDENTIFIERS

var x int

type person struct {
	first string
	last  string
}

type foo int

var y foo

func main() {
	fmt.Println("Hello")
}
