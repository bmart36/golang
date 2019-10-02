package main

import "fmt"

func main() {
	foo()
	bar("James")
	s1 := woo("Moneypenny")
	fmt.Println(s1)
	x, y := mouse("Ian", "Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

// Syntax:
// func (r receiver) identifier(parameters (return(s)) {...}
func foo() {
	fmt.Println("Hello from foo")
}

// Argument
// everything in GO is PASS BY VALUE
func bar(s string) {
	fmt.Println("Hello,",s)
}

// Return
func woo(st string) string {
	return fmt.Sprint("Hello from woo, ", st)
}

// Multiple returns
func mouse(fn, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := true
	return a, b
}