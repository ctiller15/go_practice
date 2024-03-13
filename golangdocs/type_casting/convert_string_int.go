package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "42"

	v, _ := strconv.Atoi(s)

	fmt.Println(v)

	var i int = 42
	str := strconv.Itoa(i)

	fmt.Println(str)
}
