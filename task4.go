package main

//!!!!!!!!!!!!
import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	dataChannel := make(chan string)

	var workersNumber int
	fmt.Print("Enter number of workers: ")
	fmt.Scan(&workersNumber)

	for i := 0; i < workersNumber; i++ {
		go func(i int) {
			for data := range dataChannel {
				fmt.Printf("Worker %d got: %v \n", i, data)
			}
		}(i)
	}

	go func() {
		for {
			var data string
			fmt.Println("Enter some data: ")
			fmt.Scan(&data)
			dataChannel <- data
		}
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	for signal := range c {
		fmt.Printf("Got %s, closing an application...\n", signal.String())
		close(dataChannel)
		return
	}
}
