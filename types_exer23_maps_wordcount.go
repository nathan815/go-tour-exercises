package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)
	for _, word := range words {
		if _, ok := counts[word]; ok {
			counts[word]++
		} else {
			counts[word] = 1
		}
	}
	return counts
}

func main() {
	wc.Test(WordCount)
}

