package main

import "fmt"

func main() {
	// Use a switch statement with no switch expression specified
	switch {
	case false:
		fmt.Println("This shouldn't print")

	case true:
		fmt.Println("This should print")
	}
}
