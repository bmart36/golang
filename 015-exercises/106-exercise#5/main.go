package main

import (
	"fmt"
	"math"
)

// Create a type SQUARE, CIRCLE
type square struct {
	length float64
}
type circle struct {
	radius float64
}

// Attach a method to each that calculates AREA and returns it
func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * 2
}

// Create a type SHAPE that defines an interface as anything that has the
// AREA method
type shape interface{
	area() float64
}

// Create a func INFO which takes type shape and then prints the area
func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	// Create a value of type square and another of type circle
	s1 := square {
		4,
	}
	c1 := circle {
		1.3,
	}

	// Use func info to print the area of square and circle
	info(s1)
	info(c1)
}