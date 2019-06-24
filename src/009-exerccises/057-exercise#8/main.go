package main

import "fmt"

func main() {
	// Use a switch statement with no switch expression specified
	switch {
	case (2 == 4):
		fmt.Println("This shouldn't print")

	case (2 == 2):
		fmt.Println("This should print")
	}
}
