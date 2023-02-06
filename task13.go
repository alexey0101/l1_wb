package main

import "fmt"

func main() {
	a, b := 0, 1 // initialize variables
	fmt.Printf("a: %d, b: %d\n", a, b)
	a, b = b, a // swap values
	fmt.Printf("a: %d, b: %d", a, b)
}
