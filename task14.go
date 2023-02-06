package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value interface{} //declare variable with interface type
	//value = 15
	value = "qwerty"
	//value = true
	//value = make(chan int)
	fmt.Println(reflect.TypeOf(value)) //print type of value
}
