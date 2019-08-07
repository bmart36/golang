package main

import "fmt"

func main() {
	x := map[string][]string{
		"bond_james":      {"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": {"James Bond", "Literature", "Computer Science"},
		"no_dr":           {"Being evil", "Ice cream", "Sunsets"},
	}

	x["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	// Delete a record from your map
	delete(x, "no_dr")

	// Print the map our using the "range" loop
	for k, v := range x {
		fmt.Println("Record for: ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
