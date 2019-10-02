package main

import "fmt"

func main() {
	// Build and use an anonymous func
	func() {
		for i := 0; i < 100; i++ {
			fmt.Println(i)
		}
	}()

	fmt.Println("done")
}