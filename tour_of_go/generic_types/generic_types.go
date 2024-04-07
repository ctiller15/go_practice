package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	test := List[int]{nil, 5}
	test.next = &List[int]{nil, 20}

	fmt.Println(test.next.val)
}
