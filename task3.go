package main

import "fmt"

func square(value int) int {
	return value * value
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10} //Массив
	numChan := make(chan int)

	for _, val := range arr {
		go func(val int) { // запускаем горутины, которые записывают в канал квадраты переданных в функцию значений
			numChan <- square(val)
		}(val)
	}

	result := 0
	for i := 0; i < 5; i++ {
		result += <-numChan // читаем числа с канала и суммируем их
	}

	fmt.Println("Array: ", arr)
	fmt.Printf("Sum of squares : %d", result)
}
