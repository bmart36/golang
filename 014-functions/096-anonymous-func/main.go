package main

import "fmt"

func main() {
	func(){
		fmt.Println("Anonymous func ran")
	}()

	func(x int){
		fmt.Println("The meaning of life:", x)
	}(42)
}