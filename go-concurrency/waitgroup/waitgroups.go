package main

import (
	"fmt"
	"sync" // sync package provides synchronization primitives, including WaitGroup
)

// printSomething prints a string and signals the WaitGroup when complete
// The WaitGroup pointer is passed as an argument to track completion
func printSomething(s string, wg *sync.WaitGroup) {
	// defer ensures wg.Done() is called when the function returns, even if a panic occurs
	defer wg.Done() // decrements the WaitGroup counter by 1
	fmt.Println(s)
}

func waitGroup() {
	// WaitGroup is used to wait for a collection of goroutines to finish
	// It acts as a counter of active goroutines
	// The zero value is a WaitGroup with its internal counter initialized to zero, ready to use
	var wg sync.WaitGroup

	words := []string{
		"hello",
		"world",
		"this",
		"is",
		"a",
		"test",
	}

	// Add sets the number of goroutines to wait for
	// Must be called before launching goroutines
	wg.Add(len(words))

	// Launch a goroutine for each word in the slice
	// Each goroutine processes independently and concurrently
	for i, x := range words {
		// & operator passes the WaitGroup by reference (pointer)
		// This is necessary so each goroutine references the same WaitGroup instance
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	// Wait blocks until the WaitGroup counter is zero
	// This ensures all goroutines complete before moving forward
	wg.Wait()

	// After waiting, we can reuse the same WaitGroup
	wg.Add(1)

	// This runs in the main goroutine (not concurrent)
	// But still uses the WaitGroup pattern for demonstration
	printSomething("last", &wg)
}
