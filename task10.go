package main

import (
	"fmt"
	"math"
)

func split(data []float64, step int64) map[int64][]float64 { // split data by step
	groups := map[int64][]float64{}

	for _, val := range data {
		if val >= math.MinInt64 && val <= math.MaxInt64 {
			group := int64(val) / step * step
			groups[group] = append(groups[group], val)
		}
	}

	return groups
}

func main() {
	data := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(split(data, 10))
}
