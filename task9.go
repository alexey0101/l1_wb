package main

import "fmt"

func arrayReader(arr []int) <-chan int { // function that reads array and returns channel with values
	dataChannel := make(chan int)
	go func() {
		for _, val := range arr {
			dataChannel <- val
		}
		close(dataChannel)
	}()
	return dataChannel
}

func squareNumbers(dataChannel <-chan int) <-chan int { // function that reads int's from channel and returns channel with squares
	squareChannel := make(chan int)
	go func() {
		for val := range dataChannel {
			squareChannel <- val * val
		}
		close(squareChannel)
	}()

	return squareChannel
}

func print(squareChannel <-chan int) { // function that prints values from square channel
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
