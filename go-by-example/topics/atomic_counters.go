package topics

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// The primary mechanism for managing state in Go is communication over channels.

func AtomicCountersMain() {
	// Atomic counters can be accessed by multiple goroutines
	var ops atomic.Uint64

	// A WaitGroup will help us wait for all goroutines to finish their work
	var wg sync.WaitGroup

	// Start 50 goroutines
	for range 50 {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for range 1000 {
				ops.Add(1)
			}
		}()
	}

	wg.Wait()

	// Using Load is safe to atomically read a value even while other goroutines are (atomically) updating it
	fmt.Println("ops:", ops.Load())
}
