package main

import "fmt"

func main() {
	var s string = "Hello World"
	var b []byte = []byte(s)

	fmt.Println(b)

	ss := string(b)

	fmt.Println(ss)
}
