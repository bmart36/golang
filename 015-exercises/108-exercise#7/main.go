package main

import "fmt"

func main() {
	// Assign a func to a variable, then call the func
	f := func() {
		for i := 0; i < 3; i++{
			fmt.Println(i)
		}
	}

	f()
}