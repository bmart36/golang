package main

import "fmt"

func main() {
	// Create for loop with syntax "for {}"
	bd := 1994
	for {
		if bd > 2019 {
			break
		}
		// Print out the years you've been alive
		fmt.Println(bd)
		bd++
	}
}
