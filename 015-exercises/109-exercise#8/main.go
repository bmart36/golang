package main

import "fmt"

func main() {
	// Assign the returned func to a variable
	f := foo()
	
	// Call the returned func
	f()
}

// Create a func which returns a func
func foo() func(){
	return func(){
		fmt.Println("Hallo")
	}
}