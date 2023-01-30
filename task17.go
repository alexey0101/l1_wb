package main

import "fmt"

func binarySearch(arr []int, value int) int {
	l := -1
	r := len(arr)
	for r-l > 1 {
		mid := (r + l) / 2
		if arr[mid] < value {
			l = mid
		} else {
			r = mid
		}
	}

	return r
}

func main() {
	arr := []int{1, 4, 5, 5, 5, 5, 10}
	fmt.Println(binarySearch(arr, 5))
}
