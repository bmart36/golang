package main

import (
	"fmt"
	"runtime"
)

var x int8 = -128

// var x int8 = -129

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
