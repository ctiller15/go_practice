package main

import (
	"custom_package/calculator"
	"fmt"
)

func main() {
	number1 := 9
	number2 := 5

	fmt.Println(calculator.Add(number1, number2))
	fmt.Println(calculator.Subtract(number1, number2))
}
