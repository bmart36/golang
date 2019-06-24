package main

import "fmt"

func main() {
	// In loop, print out the years you've been alive
	// in the syntax "for condition {}"
	year := 1994
	for year <= 2019 {
		fmt.Println(year)
		year++
	}
}
