package topics

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Use the built-in synchorinzation features of goroutines and channels to synchronize access to shared states across multiple goroutines.
// The channel-based approach aligns with Go's ideas of sharing memory by communicating and having each piece of data owned by exactly 1 goroutine.

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func StatefulGoroutineMain() {
	var readOpsCount uint64
	var writeOpsCount uint64

	// The channels are used by other goroutines to issue read and write requests, respectively
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		// The state is owned by a single goroutine, which guarantee that the data is never corrupted by concurrent access.
		// In order to read and write the state, other goroutines will send messages to the owning goroutine and receive corresponding reply.
		state := make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// Starts 100 goroutes to issue reads to the state-owning goroutines via the reads channel.
	for range 100 {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOpsCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Starts 10 goroutines to issue writes
	for range 10 {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool),
				}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOpsCount, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let the goroutines work for a second
	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOpsCount)
	writeOpsFinal := atomic.LoadUint64(&writeOpsCount)

	fmt.Println("readOpsFinal", readOpsFinal, ", writeOpsFinal", writeOpsFinal)
}
