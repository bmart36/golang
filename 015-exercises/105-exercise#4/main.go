package main

import "fmt"

// Create a user defined struct with:
// - the identifier "person"
// - the fields first, last and age
type person struct {
	first string
	last  string
	age   int
}

// Attach a method to type person with:
// - the identifier "speak"
// - the mothod should have the person say their name and age
func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and I'm", p.age, "years old.")
}

func main() {
	// Create a value of type person
	p1 := person{
		"James",
		"Bond",
		32,
	}
	// Call the method from the value of type person
	p1.speak()
}
