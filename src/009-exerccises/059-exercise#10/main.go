package main

import "fmt"

// Create TYPED and UNTYPED constants
const (
	x int = 42
	y     = 42
)

func main() {
	fmt.Println(x, y)
}
