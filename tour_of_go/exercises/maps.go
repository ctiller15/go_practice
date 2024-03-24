package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	answer := map[string]int{}
	for _, word := range strings.Fields(s) {
		_, ok := answer[word]
		if ok {
			answer[word]++
		} else {
			answer[word] = 1
		}
	}

	return answer
}

func main() {
	wc.Test(WordCount)
}
