package main

import (
	"fmt"
	"unicode"
)

func unique(str string) bool { // check if string has unique symbols
	m := map[rune]struct{}{}

	for _, symbol := range str {
		lower := unicode.ToLower(symbol)
		if _, ok := m[lower]; ok {
			return false
		}
		m[lower] = struct{}{}
	}

	return true
}

func main() {
	fmt.Println(unique("avbw"))
	fmt.Println(unique("aAbd"))
}
