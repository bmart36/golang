package main

import "fmt"

func main() {
	// Loop numbers between 10 and 100
	for i := 10; i <= 100; i++ {
		// Print remainder of value divided by 4
		fmt.Printf("%v/4 = %v\n", i, i%4)
	}
}
