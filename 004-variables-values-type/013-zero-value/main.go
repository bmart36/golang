package main

import "fmt"

var y string
var z int

func main() {
	// DECLARE a VARIABLE to be of a certain TYPE
	// and then ASSIGN a VALUE of that type to the variable

	fmt.Println("pritnitng the value of y", y)
	fmt.Printf("%T\n", y)

	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T\n", y)

	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
