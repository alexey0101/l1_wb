package main

import (
	"fmt"
	"unicode/utf8"
)

func reverse(str string) string {
	reversed := []rune(str)
	strLen := utf8.RuneCountInString(str)
	for i := 0; i < strLen/2; i++ {
		reversed[i], reversed[strLen-1-i] = reversed[strLen-1-i], reversed[i]
	}

	return string(reversed)
}

func main() {
	fmt.Println(reverse("qweqe"))
}
