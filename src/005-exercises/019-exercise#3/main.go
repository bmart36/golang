package main

import "fmt"

// ASSIGN the following VALUES to the three VARIABLES at 
//the package level scope:
//    a. 42 for "x"
//    b. "James Bond" for "y"
//    c. true for "z"
var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Use fmt.Sprintf to print all the VALUES to one single string
	// ASSIGN the returned value of TYPE string using short declaration
	// operator to a VARIABLE with the IDENTIFIER "s"
	s := fmt.Sprintf("%v, %v, %v", x, y, z)

	// Print out the value stored by variable "s"
	fmt.Println(s)
}