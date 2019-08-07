package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Using APPEND and SLICING, get the following values ASIGNED to the
	// variable "y": [42, 43, 44, 48, 49, 50, 51]
	y := append(x[:3], x[6:]...)

	// Print out slice
	fmt.Println(y)
}
