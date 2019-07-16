package main

import "fmt"

func main() {
	// Use a switch statement with the switch expression specified
	// as a variable of TYPE string wiht the IDENTIFIER "favSport"
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("go to the pool!")
	case "surfing":
		fmt.Println("go to hawaii")
	case "wingsuit flying":
		fmt.Println("what would oyu like me to say at your funeral")

	}
}
