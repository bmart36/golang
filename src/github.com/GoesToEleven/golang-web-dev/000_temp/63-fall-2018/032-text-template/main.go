package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt", "two.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}

	fmt.Println("\n-----")
	tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)

	fmt.Println("\n\n-----")
	tpl.ExecuteTemplate(os.Stdout, "two.txt", "James")
	fmt.Println("\n\n")
}
