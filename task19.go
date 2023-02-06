package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(str string) string { // reverse string
	reversed := []rune(str)
	strLen := utf8.RuneCountInString(str) // get string length in runes
	for i := 0; i < strLen/2; i++ {       // swap runes
		reversed[i], reversed[strLen-1-i] = reversed[strLen-1-i], reversed[i]
	}

	return string(reversed)
}

func main() {
	fmt.Println(reverse("qweqe"))
}
