package main

import "fmt"

func main() {
	fmt.Println("bla bla bla")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("In foo")
}

func bar() {
	fmt.Println("end")
}

// control flow:
// (1) sequence
// (2) loop: iterative
// (3) conditional
