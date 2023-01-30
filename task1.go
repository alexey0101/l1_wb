package main

import "fmt"

type Human struct {
	age    int
	height int
	weight int
}

func (human Human) SayHi() {
	fmt.Println("hi")
}

type Action struct {
	name string
	Human
}

func main() {
	action := Action{name: "Greeting"}
	action.SayHi()
}
