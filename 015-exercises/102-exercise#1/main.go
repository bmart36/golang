package main

import "fmt"

func main() {
	// Call both funcs
	x := foo()
	y, z := bar()

	// Print ouy all their results
	fmt.Println(x, y, z)
}

// Create a func with the identifier foo that returns an int
func foo() int{
	return 42
}

// Create a func with the identifier bar that returns an int and a string
func bar() (int, string){
	return 1984, "Big Brither is Watching"
}