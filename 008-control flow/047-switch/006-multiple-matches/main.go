package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny", "Bond", "Dr No":
		fmt.Println("miss money or bond or dr no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("This should not print2")
	}
}
