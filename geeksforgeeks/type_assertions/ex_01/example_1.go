// Golang program to illustrate
// the concept of type assertions
package main

import "fmt"

func main() {
	// an interface that has
	// a string value
	var value interface{} = "GeeksforGeeks"

	// retrieving a value
	// of type string and assigning
	// it to value1 variable
	var value1 string = value.(string)

	// printing the concrete value
	fmt.Println(value1)

	// this will panic as interface
	// does not have int type
	var value2 int = value.(int)

	fmt.Println(value2)
}
