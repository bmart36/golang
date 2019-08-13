package main

import "fmt"

// Create new type vehicule, underlying type struct,
// with fields doors, color
type vehicule struct {
	doors int
	color string
}

// Create two new types: truck & sedan, underlying type struct
// and embed the "vehicule" type in both
type truck struct {
	vehicule
	fourWheel bool
}

type sedan struct {
	vehicule
	luxury bool
}

func main() {
	// Using a composite literal, create a value fo type truck and assign
	// values to the fields
	t := truck{
		vehicule: vehicule{
			doors: 2,
			color: "blue",
		},
		fourWheel: true,
	}

	// Using a composite literal, create a value fo type sedan and assign
	// values to the fields
	s := sedan{
		vehicule: vehicule{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}

	// Print out each of these values
	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.doors)
	fmt.Println(s.doors)
}
