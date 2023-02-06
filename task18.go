package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
}

func (c *Counter) Increment() { //concurrent safe increment
	atomic.AddInt32(&c.value, 1)
}

func (c *Counter) Value() int32 { //concurrent safe get value
	return atomic.LoadInt32(&c.value)
}

func main() {
	counter := Counter{0}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() { //10 goroutines for incrementing counter
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()                    //wait for all goroutines to finish
	fmt.Println(counter.Value()) //10
}
