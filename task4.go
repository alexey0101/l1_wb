package main

//!!!!!!!!!!!!
import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	dataChannel := make(chan string) // channel for data

	var workersNumber int
	fmt.Print("Enter number of workers: ") // number of workers
	fmt.Scan(&workersNumber)               // read number of workers

	for i := 0; i < workersNumber; i++ {
		go func(i int) { // start workers
			for data := range dataChannel {
				fmt.Printf("Worker %d got: %v \n", i, data)
			}
		}(i)
	}

	go func() {
		for { // read data from console
			var data string
			fmt.Println("Enter some data: ")
			fmt.Scan(&data)
			dataChannel <- data
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt) // listen for Ctrl+C
	for signal := range c {
		fmt.Printf("Got %s, closing an application...\n", signal.String())
		close(dataChannel)
		return
	}
}
