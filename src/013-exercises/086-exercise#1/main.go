package main

import "fmt"

// Create your own type "person" which will have an underlying type "struct"
// so that it can store a first name, last name and favorite ice cream flavors
type person struct {
	first      string
	last       string
	favFlavors []string
}

func main() {
	// Create two VALUES of TYPE person
	p1 := person{
		first: "James",
		last:  "Bond",
		favFlavors: []string{
			"chocolate",
			"martini",
			"rum and coke",
		},
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		favFlavors: []string{
			"strawberry",
			"vanilla",
			"capuccino",
		},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	// Print out values ranging over the elements in the slice which
	// stores the favorite flavors
	for i, v := range p1.favFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFlavors {
		fmt.Println(i, v)
	}
}
