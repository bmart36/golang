package main

import "fmt"

func main() {

	ii := []int{1,2,3,4,5,6,7,8,9}
	// pass in a value of type []int into the func foo (unfurl the []int)
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{1,2,3,4,5,6,7,8,9}
	n2 := bar(ii2)
	fmt.Println(n2)
}

// Create func with identifier foo that:
// - takes in a variadic parameter of type int
// - returns the sum of the values of type int passed
func foo(xi ...int) int{
	sum := 0
	for _, v := range xi{
		sum += v
	}
	return sum
}

// Create a func with the indetifier bar that:
// - takes in a parameter of type []int
// - returns the sum of all the values of type int passed
func bar(xi []int) int{
	sum := 0
	for _, v := range xi{
		sum += v
	}
	return sum
}