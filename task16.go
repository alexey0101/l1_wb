package main

import "fmt"

func quicksort(arr []int, l, r int) { //quick sort
	if l < r { //if left index less than right index
		mid := arr[(l+r)/2] //mid value
		i := l
		j := r
		for i < j { //while left index less than right index
			for arr[i] < mid { //if value less than mid value, then increase left index
				i++
			}
			for arr[j] > mid { //if value more than mid value, then decrease right index
				j--
			}
			if i < j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		quicksort(arr, l, j)   // call quicksort for left part
		quicksort(arr, j+1, r) // call quicksort for right part
	}
}

func main() {
	arr := []int{1, 2, 5, 0, -1}
	quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
