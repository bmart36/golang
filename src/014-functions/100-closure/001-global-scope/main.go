package main

import "fmt"

var x int

func main() {

	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
}

func foo() {
	x++
	fmt.Println("hello")
}