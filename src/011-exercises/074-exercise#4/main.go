package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Append the value 52 to the slice
	x = append(x, 52)

	// Print out the slice
	fmt.Println(x)

	// In ONE STATEMENT append to the slice the values 53, 54 and 55
	x = append(x, 53, 54, 55)

	// Print out the slice
	fmt.Println(x)

	// Append the following slice: y := []int{56, 57, 58, 59, 60}
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	// Print out the slice
	fmt.Println(x)
}
