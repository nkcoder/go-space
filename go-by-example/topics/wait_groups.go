package topics

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup waits for a collection of goroutines to finish.
// A WaitGroup must not be copied after first use. So if a WaitGroup is explicitly passed into functions, it should be done by pointer.
func wgWorker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func WaitGroupsMain() {
	var wg sync.WaitGroup

	// Launch several goroutines and increment the counter for each.
	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			// Decrements the counter by 1
			defer wg.Done()
			wgWorker(i)
		}()
	}

	// Block untile the counter goes back to 0: all the workers notified they're done.
	wg.Wait()
}
