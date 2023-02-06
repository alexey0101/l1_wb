package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string { // reverse words in string
	words := strings.Fields(str) // split string to words
	wordsLen := len(words)

	for i := 0; i < wordsLen/2; i++ { // swap words
		words[i], words[wordsLen-1-i] = words[wordsLen-1-i], words[i]
	}

	return strings.Join(words, " ") // join words to string with space
}

func main() {
	str := "qwe dc ns"
	fmt.Println(reverseWords(str))
}
