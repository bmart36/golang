package main

import "fmt"

func main() {
	f1 := foo()
	fmt.Println(f1())
	fmt.Println(f1())
	f2 := foo()
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
	fmt.Println(f2())
}

// Syntax:
// func (r receiver) identifier(parameters (return(s)) {...}
func foo() func() int{
	x := 0
	return func() int {
		x++
		return x
	}
}