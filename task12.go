package main

import "fmt"

type Set struct {
	Values map[string]struct{}
}

func NewSet(values []string) *Set {
	set := Set{map[string]struct{}{}}
	for _, val := range values {
		set.Values[val] = struct{}{}
	}
	return &set
}

func main() {
	set := NewSet([]string{"cat", "cat", "dog", "cat", "tree"})
	fmt.Println(set.Values)
}
