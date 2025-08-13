package topics

import (
	"fmt"
	"time"
)

/* channels are pipes that connect concurrent goroutines, you can send messages into channels from one goroutine and receive those messages from another goroutine
 */

func unbufferedChannel() {
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
}

func bufferedChannel() {
	// buffered channels accept a limited number of values without a corresponding receiver for thoese values
	bufferedMessages := make(chan string, 2)
	// because the channel is buffered, we can send those values into the channel without a corresponding concurrent receive
	bufferedMessages <- "buffered"
	bufferedMessages <- "channel"

	fmt.Println(<-bufferedMessages)
	fmt.Println(<-bufferedMessages)
}

// this function will run in a goroutine
// the done channel will be used to notify another goroutine that its work is done.
func channelWorker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func workerNotify() {
	// start a working goroutine, give it the channel to nofify on
	done := make(chan bool, 1)
	go channelWorker(done)

	// block until we receive a notification from the worker on the channel
	<-done
}

// Closing a channel indicates that no more values will be sent on it, this can be useful to communicate completion to the channel's receivers.
func closingChannel() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	// Reading from a close channel succeeds immediately, returning the zero value of the underlying type.
	// The optional second return value is true if the value received was delivered by a successful send operation to the channel,
	// or false if it was a zero value generated because the channel is closed and empty.
	_, ok := <-jobs
	fmt.Println("received more jobs after close", ok)
}

func rangeOverChannel() {
	queue := make(chan string, 2)
	queue <- "hello"
	queue <- "world"

	close(queue)

	// The range iterates over each element as it's received from queue.
	// The iteration terminates after receiving 2 elements since we closed the channel.
	for m := range queue {
		fmt.Println(m)
	}
}

func ChannelMain() {
	unbufferedChannel()

	bufferedChannel()

	workerNotify()

	closingChannel()

	rangeOverChannel()
}
