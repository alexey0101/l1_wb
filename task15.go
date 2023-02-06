package main

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	justString = string(string([]rune(v)[:100])) // create string from slice of runes
}

func main() {
	someFunc()
}
