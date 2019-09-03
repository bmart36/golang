package main

import "fmt"

func main() {
	// Assing int to a variable
	x := 42

	// Print int in decimal, binary, and hex
	fmt.Printf("%d, %b, %#x\n", x, x, x)

	// Shift bits of int over 1 position to the left and
	// assign it to a variable
	y := x << 1

	// Print new variable in decimal, binary, and hex
	fmt.Printf("%d, %b, %#x\n", y, y, y)
}
