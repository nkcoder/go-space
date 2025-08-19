package topics

import (
	"fmt"
	"time"
)

// a goroutine is a lightweight thread of execution
func fg(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func GoroutineMain() {
	fg("direct")

	// to invoke a function in a goroutine, use `go f()`, the new goroutine will execute concurrently with the calling one
	go fg("goroutine")

	// you can also start a goroutine for an anonymous function call
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// our two function calls are running asynchronously in separate goroutines now, wait for them to finish
	// for a more robust approach, use a WaitGroup
	time.Sleep(time.Second)

	// we'll see the output of the blocking call first, then the output of the two goroutines.
	// the goroutines' output may be interleaved, because goroutines are being run concurrently by the Go runtime.
	fmt.Println("done")
}
