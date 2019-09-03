package main

import "fmt"

// Create own type with the underlying type being int
type myType int

// Create a VARIABLE of the new TYPE with the IDENTIFIER
// "x" using the "VAR" keyword
var x myType

func main() {
	// Print out the value of variable "x"
	fmt.Println(x)

	// Print out the type of variable "x"
	fmt.Printf("%T\n", x)

	// ASSIGN 42 to the VARIABLE "x" using the "=" OPERATOR
	x = 42

	// Print out the value of the variable "x"
	fmt.Println(x)
}