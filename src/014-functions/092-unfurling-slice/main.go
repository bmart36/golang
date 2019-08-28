package main

import "fmt"

func main() {
	xi := []int{2,3,4,5,6,7,8,9}
	// Passing individual values in a slice by using the "..." operator
	x := sum(xi...)
	fmt.Println("The total is,", x)

	// Passsing zero values is possible
	x1 := sum()
	fmt.Println("The total is,", x1)

	x2 := sumAlt("The total is,", 4, 5, 6)
	fmt.Println(x2)
}

func sum(x ...int) int{
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position,", i,"we are now adding,", v, "to the total which is now", sum)
	}
	return sum
}

// Variadic parameter should be last parameter e.g. func foo(s string, x ...int)
// fun foo(x ..int, s string) is not allowed
func sumAlt(s string, x ...int) string{
	sum := 0
	for _, v := range x {
		sum += v
	}
	return fmt.Sprint(s, " ", sum)
}