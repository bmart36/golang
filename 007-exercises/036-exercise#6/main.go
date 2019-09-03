package main

import "fmt"

// Using iota, create 4 constants for the NEXT 4 years
const (
	_     = iota
	next1 = 2019 + iota
	next2 = 2019 + iota
	next3 = 2019 + iota
	next4 = 2019 + iota
)

func main() {
	fmt.Println(next1, next2, next3, next4)
}
