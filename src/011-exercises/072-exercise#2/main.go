package main

import "fmt"

func main() {
	// Using a COMPOSITE LITERAL create a SLICE of TYPE int and assing 10 VALUES
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// Range over the slice and print the values out
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Using format printing, print out the TYE of the array
	fmt.Printf("%T\n", x)
}
