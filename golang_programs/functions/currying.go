package main

import "fmt"

func squareSum(x int) func(int) func(int) int {
	return func(y int) func(int) int {
		return func(z int) int {
			return x*x + y*y + z*z
		}
	}
}

func main() {
	// 5*5 + 6*6 + 7*7
	fmt.Println(squareSum(5)(6)(7))
}
