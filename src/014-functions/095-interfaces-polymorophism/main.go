package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last, " - the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I am", p.first, p.last, " - the person speak")
}

// Interfaces allow to define behaviour  and to do polymorphism
type human interface{
	// Any type who has the method spead is also of type human
	speak()
}

func bar(h human) {
	// Assertion
	switch h.(type){
	case person:
		fmt.Println(" I was passed into barrrr", h.(person).first)
	case secretAgent:
		fmt.Println(" I was passed into barrrr", h.(secretAgent).first)
	}
	fmt.Println(" I was passed into bar", h)
}

type hotdog int

func main() {
	// Values can be of more than one type e.g. sa1 is of type secretAgent
	// but since it has the method speak it is alos of type human
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Miss",
			"Moneypenny",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last: "yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)
	p1.speak()

	bar(sa1)
	bar(sa2)
	bar(p1)

	//conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}