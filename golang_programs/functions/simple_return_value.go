package main

import "fmt"

// function with int as return type
func add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	// Accepting return value in variable
	sum := add(20, 30)
	fmt.Println(sum)
}
