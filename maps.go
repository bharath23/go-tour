/* Exercise: Maps */
package main

import (
	"code.google.com/p/go-tour/wc"
	"string"
)

func WordCount(s string) map[string]int {
	count := make(map[string]int)
	for _, word := range strings.Fields(s) {
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
