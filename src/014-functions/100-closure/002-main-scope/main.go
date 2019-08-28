package main

import "fmt"


func main() {

	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println("hello")
}