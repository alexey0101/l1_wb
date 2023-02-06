package main

import (
	"fmt"
	"strings"
)

type Set struct {
	values map[int]struct{}
}

func NewSet(values []int) *Set {
	valuesMap := map[int]struct{}{}
	for _, val := range values {
		valuesMap[val] = struct{}{} // fill values
	}
	return &Set{valuesMap}
}

func (baseSet *Set) Intersection(set *Set) *Set { // intersection of two sets
	values := []int{}

	var iterateSet, biggestSet *Set
	if len(baseSet.values) < len(set.values) { // choose set with less values
		iterateSet = baseSet
		biggestSet = set
	} else {
		iterateSet = set
		biggestSet = baseSet
	}

	for key, _ := range iterateSet.values { // iterate over set with less values
		if _, ok := biggestSet.values[key]; ok { // check if value exists in another set
			values = append(values, key) // add value to result set
		}
	}

	return NewSet(values)
}

func (set *Set) String() string { // print set
	var b strings.Builder
	for key, _ := range set.values {
		fmt.Fprintf(&b, "%d ", key)
	}

	return strings.TrimSpace(b.String())
}

func main() {
	set1 := NewSet([]int{1, 4, 6, 7, 8, 10, 15, 20, 22, 23})
	set2 := NewSet([]int{0, 5, 7, 11, 12, 15, 20, 21})

	fmt.Println(set1.Intersection(set2))
}
