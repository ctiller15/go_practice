package main

import "fmt"

// The continue statement skips one or more loop iterations.
func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
