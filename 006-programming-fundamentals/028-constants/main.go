package main

import "fmt"

const (
	// UYPED constant
	a int = 42
	// UNTYPED constants
	b = 42.78
	c = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
