package main

import "fmt"

func arrayReader(arr []int) <-chan int {
	dataChannel := make(chan int)
	go func() {
		for _, val := range arr {
			dataChannel <- val
		}
		close(dataChannel)
	}()
	return dataChannel
}

func squareNumbers(dataChannel <-chan int) <-chan int {
	squareChannel := make(chan int)
	go func() {
		for val := range dataChannel {
			squareChannel <- val * val
		}
		close(squareChannel)
	}()

	return squareChannel
}

func print(squareChannel <-chan int) {
	for val := range squareChannel {
		fmt.Println(val)
	}
}

func main() {
	arr := []int{1, 4, 6, 8, 10, 11, 15}
	dataChannel := arrayReader(arr)
	squareChannel := squareNumbers(dataChannel)
	print(squareChannel)
}
