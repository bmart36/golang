package main

import "fmt"

func main() {
	// Create a slice to store the names of all the states in the USA
	x := make([]string, 50, 50)

	// Length & capacity of slice
	fmt.Println(len(x)) // 50
	fmt.Println(cap(x)) // 50

	x = []string{
		" Alabama", " Alaska", " Arizona", " Arkansas", " California",
		" Colorado", " Connecticut", " Delaware", " Florida", " Georgia",
		" Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas",
		" Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts",
		" Michigan", " Minnesota", " Mississippi", " Missouri", " Montana",
		" Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico",
		" New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma",
		" Oregon", " Pennsylvania", " Rhode Island", " South Carolina",
		" South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia",
		" Washington", " West Virginia", " Wisconsin", " Wyoming"}

	// Print out all values, along with their index position in the slice, without using the range clause
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
