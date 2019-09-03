package main

import "fmt"

func main() {
	// In loop, print out the years you've been alive
	// in the syntax "for condition {}"
	bd := 1994
	for bd <= 2019 {
		fmt.Println(bd)
		bd++
	}
}
