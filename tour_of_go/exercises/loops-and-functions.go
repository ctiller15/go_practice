package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	eps := x
	for math.Abs(eps) > .0000000001 {
		eps = (z*z - x) / (2 * z)
		z -= eps
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
}
