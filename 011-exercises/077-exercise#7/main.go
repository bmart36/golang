package main

import "fmt"

func main() {
	// Create a slice of string ([][]string) and store th following data in the
	// multi-dimensional slice:
	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	x := [][]string{a, b}

	// Range over the records, then range over the data in each record
	for i, xs := range x {
		fmt.Println("record: ", i)
		for j, v := range xs {
			fmt.Printf("\t index position %v \t value: %v \n", j, v)
		}
	}
}
