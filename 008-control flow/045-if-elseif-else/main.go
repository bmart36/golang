package main

import "fmt"

func main() {
	x := 434
	if x == 40 {
		fmt.Println("The value is 40")
	} else if x == 41 {
		fmt.Println("The value is 41")
	} else if x == 42 {
		fmt.Println("The value is 42")
	} else if x == 43 {
		fmt.Println("The value is 43")
	} else {
		fmt.Println("The value is not 40, 41, 42 or 43")
	}
}
