package main

import "fmt"

// DECLARE three VARIABLES using var. They should have package level
// scope. No VALUES should be assigned to the variables. Using the following
// IDENTIFIERS for the variables, make sure they are of the correct TYPE
//     a. identifier "x" type int
//     b. identifier "y" type string
//     c. identifier "x" type bool
var x int
var y string
var z bool

func main() {
	// Print out the values for each identifier
	fmt.Println("x:",x)
	fmt.Println("y:",y)
	fmt.Println("z:",z)

	// The values assigned to the unassigned variables are:
	//     a. 0 for "x"
	//     b. "" for "y"
	//     c. false for "z"
}