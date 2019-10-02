package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello")
}

// Use the "defer" keyword to show that a deferred func runs after the func
// containing it exits
func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}