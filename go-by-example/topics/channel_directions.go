package topics

import "fmt"

/*
* When using channels as function parameters, you can specify if the channel is meant to only send or receive values.
* This increases type-safety.
 */

// Accepts send-only channel, it would be a compile-time error to try to receive on this channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// Accepts one channel for receives and the other for sends
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirectionMain() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "processed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
