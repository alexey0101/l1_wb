package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{}
	//value = 15
	value = "qwerty"
	//value = true
	//value = make(chan int)
	fmt.Println(reflect.TypeOf(value))
}
