package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println(time.Now())
	sleep(time.Duration(5 * time.Second))
	fmt.Println(time.Now())
}
