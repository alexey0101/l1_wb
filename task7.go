package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Map struct {
	values map[string]int
	lock   sync.RWMutex //read-write mutex
}

func (m *Map) Set(key string, value int) { //concurent safe set
	m.lock.Lock()
	m.values[key] = value
	m.lock.Unlock()
}

func (m *Map) Get(key string) (int, bool) { //concurent safe get
	m.lock.RLock()
	defer m.lock.RUnlock()
	value, ok := m.values[key]
	return value, ok
}

func main() {
	m := Map{values: map[string]int{}}

	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ { //20 goroutines for setting/getting values
		go func() {
			defer wg.Done()
			i := rand.Intn(10)
			fmt.Printf("Setting value %d in map\n", i)
			m.Set(string(i), i)
			fmt.Println("Done\n==========================")
		}()
		go func() {
			defer wg.Done()
			i := rand.Intn(10)
			fmt.Printf("Getting value by key %d\n", i)
			value, ok := m.Get(string(i))
			fmt.Printf("val = %d, exists = %v\n", value, ok)
			fmt.Println("Done\n==========================")
		}()
	}

	wg.Wait()
}
