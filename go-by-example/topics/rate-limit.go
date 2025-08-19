package topics

import (
	"fmt"
	"time"
)

// We can use a unbuffered channel to achieve the rate limiting
func limiter() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// The limiter will receive a value every 200ms, it is the regulator in our rate limiting scheme
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking on a receive from the limiter channel before serving each request, we limit to 1 request every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}
}

// We may want to allow short bursts of requests in the rate limiting while preserving the overall rate limit.
// We can accomplish this by buffering our limiter channel(using buffered channel).
func burstyLimiter() {
	// Fill up the channel to represent allowed bursting
	burstyLimiter := make(chan time.Time, 3)
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// Try to add a new value to the limter every 200ms, up to its limit 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Simulate 5 more incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	// We serve the first 3 requests immediately because of the burstable rate limiting
	// Then serve the remaining 2 with 200ms delays each.
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func RateLimitMain() {
	limiter()

	fmt.Println("------------")

	burstyLimiter()
}
