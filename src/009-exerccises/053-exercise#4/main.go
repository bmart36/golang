package main

import "fmt"

func main() {
	// Create for loop with syntax "for {}"
	year := 1994
	for {
		if year > 2019 {
			break
		}
		// Print out the years you've been alive
		fmt.Println(year)
		year++
	}
}
