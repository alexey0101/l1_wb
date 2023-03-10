package main

import (
	"fmt"
	"math"
	"math/big"
)

func sum(a, b *big.Int) *big.Int { // sum of two big numbers
	return a.Add(a, b)
}

func multiply(a, b *big.Int) *big.Int { // multiply of two big numbers
	return a.Mul(a, b)
}

func division(a, b *big.Int) *big.Int { // division of two big numbers
	return a.Div(a, b)
}

func substraction(a, b *big.Int) *big.Int { // substraction of two big numbers
	return a.Sub(a, b)
}

func main() {
	a, b := big.NewInt(int64(math.Pow(2, 21))), big.NewInt(int64(math.Pow(2, 20))) // initialize variables
	fmt.Printf("a = %v, b = %v\n", a, b)
	fmt.Printf("Sum: %v\n", sum(a, b))
	fmt.Printf("Multiply: %v\n", multiply(a, b))
	fmt.Printf("Division (a/b): %v\n", division(a, b))
	fmt.Printf("Substraction (a - b): %v\n", substraction(a, b))
}
