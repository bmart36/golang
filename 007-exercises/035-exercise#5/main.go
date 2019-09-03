package main

import "fmt"

func main() {
	// Create variable of type string using raw string literal
	x := `Hello
	world
	"This is
	a raw string"`
	fmt.Println(x)
}
