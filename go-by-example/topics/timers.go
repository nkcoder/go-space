package topics

import (
	"fmt"
	"time"
)

// Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified
// at that time.

func TimerMain() {
	// Creates a new timer that will send the current time on its channel after the duration
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks on the timer's channel until it sends a value indicating that the timer fired.
	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()

	// You can cancel a timer before it fires.
	// It returns true if the call stops the timer, returns false if the timer is already expired or stoped.
	if stop2 := timer2.Stop(); stop2 {
		fmt.Println("Timer 2 stopped")
	}

	// Give enough time for the timers to fire or stop
	time.Sleep(2 * time.Second)
}

// Timers are for when you want to do something once in the future, Tickers are for when you want to do something repeatedly at regular intervals.

func TickerMain() {
	// The ticker contains a channel that will send the current time on the channel after each tick
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("done")
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Tickers can be stopped like timers. Once it is stopped, it won't receive any more values on the channel.
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")
}
