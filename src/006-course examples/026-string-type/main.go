package main

import "fmt"

func main() {
	// Strings in go are immutable and slices of bytes
	s := "Hello, is it me you're looking for"
	fmt.Println(s)
	s = "Hello there"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Print utf-8 code point
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}