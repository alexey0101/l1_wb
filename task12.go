package main

import "fmt"

type Set struct {
	Values map[string]struct{}
}

func NewSet(values []string) *Set {
	set := Set{map[string]struct{}{}} // create empty set
	for _, val := range values {
		set.Values[val] = struct{}{} // fill values from slice
	}
	return &set
}

func main() {
	set := NewSet([]string{"cat", "cat", "dog", "cat", "tree"}) // create set from slice
	fmt.Println(set.Values)
}
