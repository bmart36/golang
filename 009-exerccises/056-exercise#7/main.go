package main

import "fmt"

// Create program that uses "else if" and "else"
func main() {
	x := 5
	y := 9
	if x == y {
		fmt.Println("x and y are equal")
	} else if x > y {
		fmt.Println("x is bigger than y")
	} else {
		fmt.Println("x is smaller than y")
	}
}
