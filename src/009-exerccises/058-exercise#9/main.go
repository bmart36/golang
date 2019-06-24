package main

import "fmt"

func main() {
	// Use a switch statement with the switch expression specified
	// as a variable of TYPE string wiht the IDENTIFIER "favSport"
	favSport := "basketball"
	switch favSport {
	case "football":
		fmt.Println("My favorite sport is football")

	case "basketball":
		fmt.Println("My favorite sport is basketball")
	}
}
