package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	value int32
}

func (c *Counter) Increment() {
	atomic.AddInt32(&c.value, 1)
}

func (c *Counter) Value() int32 {
	return atomic.LoadInt32(&c.value)
}

func main() {
	counter := Counter{0}

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.Value())
}
