package main

import "fmt"

func main() {
	// Pass a func into a func as an argument
	foo(sqr, 4)
}

func sqr(x int) int{
	return x*x
}

func foo(f func(x int) int, xi int) {
	fmt.Println(f(xi))
}