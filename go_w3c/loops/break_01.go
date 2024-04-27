package main

import "fmt"

// break terminates the loop execution.
func main() {
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
