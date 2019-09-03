package main

import "fmt"

func main() {
	// Create and use an anonymous struct
	anon := struct {
		name      string
		friends   map[string]int
		favDrinks []string
	}{
		name: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(anon.name)
	fmt.Println(anon.friends)
	fmt.Println(anon.favDrinks)

	for k, v := range anon.friends {
		fmt.Println(k, v)
	}

	for i, v := range anon.favDrinks {
		fmt.Println(i, v)
	}
}
