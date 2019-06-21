package main

import "fmt"

func main() {
	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is Q")
	default:
		fmt.Println("This should not print2")
	}
}
