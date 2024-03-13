package main

import "fmt"

func main() {
	var totalsum int = 846
	var number int = 19
	var avg float32

	avg = float32(totalsum) / float32(number)

	fmt.Printf("Average = %f\n", avg)
}
