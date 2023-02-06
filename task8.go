package main

import "fmt"

func setBit(value *int64, num int, bit bool) { //set bit in value by num position
	var rotatedBit int64
	rotatedBit = 1 << num

	if bit {
		*value = *value | rotatedBit
	} else {
		*value = *value & rotatedBit
	}

}

func main() {
	var value int64
	value = 1
	setBit(&value, 2, true)
	fmt.Println(value)

}
