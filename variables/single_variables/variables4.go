package main

import (
	"fmt"
)

// Declaring variables outside of a function with var.

var a int
var b int = 2
var c = 3

func main() {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
