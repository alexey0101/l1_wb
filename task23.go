package main

import "fmt"

func deleteElement(slice []int, idx int) []int { //delete element from slice by index
	return append(slice[:idx], slice[idx+1:]...)
}

func main() {
	arr := []int{1, 3, 5, 7, 10}
	fmt.Println(arr)
	arr = deleteElement(arr, 2)
	fmt.Println(arr)
}
