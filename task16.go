package main

import "fmt"

func quicksort(arr []int, l, r int) {
	if l < r {
		mid := arr[(l+r)/2]
		i := l
		j := r
		for i < j {
			for arr[i] < mid {
				i++
			}
			for arr[j] > mid {
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		quicksort(arr, l, j)
		quicksort(arr, j+1, r)
	}
}

func main() {
	arr := []int{1, 2, 5, 0, -1}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
