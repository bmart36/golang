package main

import "fmt"

func main() {
	// Create a map with a KEY of type string which is a person's "last_first" name
	// and a value of TYPE []string which stores their favorite things. Store three
	// records in map

	// "James", "Bond", "Shaken, not stirred"
	// "Miss", "Moneypenny", "Helloooooo, James."
	x := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	// Print out all of the values, along with their index position in the slice
	for k, v := range x {
		fmt.Println("Record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
