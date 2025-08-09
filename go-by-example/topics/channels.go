package topics

import (
	"fmt"
	"time"
)

/* channels are pipes that connect concurrent goroutines, you can send messages into channels from one goroutine and receive those messages from another goroutine
 */

// this function will run in a goroutine
// the done channel will be used to notify another goroutine that its work is done.
func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func ChannelMain() {
	// create a new channel with `make(chan val-type)`, channels are typed by the values they convey
	// by default channels are unbuffered which means sends and recevies block until both the sender and receiver are ready
	unbufferedMessages := make(chan string)

	// send a value into a channel using the `channel <-` syntax
	// if we don't put it inside a goroutine, it will block forever because there is no receiver ready yet
	go func() {
		unbufferedMessages <- "ping"
	}()

	// the `<-channel` syntax recevies a value from the channel
	msg := <-unbufferedMessages
	fmt.Println(msg)

	// buffered channels accept a limited number of values without a corresponding receiver for thoese values
	bufferedMessages := make(chan string, 2)
	// because the channel is buffered, we can send those values into the channel without a corresponding concurrent receive
	bufferedMessages <- "buffered"
	bufferedMessages <- "channel"

	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)

	// start a working goroutine, give it the channel to nofify on
	done := make(chan bool, 1)
	go worker(done)

	// block until we receive a notification from the worker on the channel
	<-done
}
