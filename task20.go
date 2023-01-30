package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {
	words := strings.Fields(str)
	wordsLen := len(words)

	for i := 0; i < wordsLen/2; i++ {
		words[i], words[wordsLen-1-i] = words[wordsLen-1-i], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	str := "qwe dc ns"
	fmt.Println(reverseWords(str))
}
