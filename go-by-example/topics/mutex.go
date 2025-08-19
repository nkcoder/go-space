package topics

import (
	"fmt"
	"sync"
)

// For more complex state we can use a mutex to safely access data across multiple goroutines.
// Note that mutexes must not be copied, so we should pass it by pointer.

type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// We want to update the map in the Container concurrently from multiple goroutines, we add a mutex too synchronize access.
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func MutexMain() {
	// The zero vaue of a mutex is usable as-is, so no initialization is required here.
	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup

	doIncrement := func(name string, n int) {
		for range n {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 10000)
	go doIncrement("b", 20000)
	go doIncrement("c", 30000)

	wg.Wait()

	fmt.Println(c.counters)
}
