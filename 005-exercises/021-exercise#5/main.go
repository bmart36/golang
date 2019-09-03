package main

import "fmt"

type myType int
var x myType

// In the package level scope, using the "var" keyword, create
// a VARIABLE with the IDENTIFIER "y". The variable should be 
// of the UNDERLYING TYPE of the custom TYPE "x"
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	// Use CONVERSION to convert the TYPE of the VALUE stored in "x"
	// to the UNDERLYING TYPE, use the "=" operator to ASSIGN that
	// value to "y"
	y = int(x)

	// Print out the value stored in "y"
	fmt.Println(y)

	// Print out the type of "y"
	fmt.Printf("%T\n", x)
}