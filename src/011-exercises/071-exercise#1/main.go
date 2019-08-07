package main

import "fmt"

func main() {
	// Using a COMPOSITE LITERAL create an ARRAY which holds 5 VALUES
	// of TYPE int and assing VALUES to each index position
	x := [5]int{42, 43, 44, 45, 46}

	// Range over the array and print the values out
	for i, v := range x {
		fmt.Println(i, v)
	}

	// Using format printing, print out the TYE of the array
	fmt.Printf("%T\n", x)
}
