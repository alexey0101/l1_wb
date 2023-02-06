package main

import (
	"context"
	"fmt"
	"time"
)

func reader(context context.Context) <-chan int {
	data := []int{2, 3, 5, 7, 11, 13, 17, 19}
	dataChannel := make(chan int) // channel for data

	go func() {
		for _, val := range data { // read data from slice with timeout
			select {
			case <-context.Done():
				close(dataChannel)
				return
			default:
				dataChannel <- val
			}
		}
		close(dataChannel)
		return
	}()

	return dataChannel
}

func writer(dataChannel <-chan int, context context.Context) {
	for val := range dataChannel { // write data to console from data channel with timeout
		select {
		case <-context.Done():
			return
		default:
			fmt.Println(val)
		}
	}
}

func main() {
	var timeout int
	fmt.Print("Enter app timeout: ")
	fmt.Scan(&timeout)

	ctx, _ := context.WithTimeout(context.Background(), time.Duration(time.Second*time.Duration(timeout))) // context with timeout
	dataChannel := reader(ctx)
	writer(dataChannel, ctx)
}
