// Printing multiple variables via one Print() statement

package main

import (
	"fmt"
)

func main() {
	var i, j string = "Hello", "World"

	fmt.Print(i, "\n", j)
	fmt.Print("\n")
	fmt.Print(i, " ", j) // Using a space instead.
}
