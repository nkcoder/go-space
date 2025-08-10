package topics

import (
	"fmt"
	"time"
)

// `select` is a control structure used for handling multiple channel operations concurrently.
// It functions similarly to `switch` statement but it's specifically designed for communication actions on channels, such as sending or receiving data.
// It blocks until one of its cases, representing a channel operation, can proceed. If multiple channel operations are ready simultaneously, it will randomly choose one to execute.
// An optional `default` clause can be included. If no channel operation is immediately ready, the `default` branch will be executed, preventing the `select` statement from blocking.

func selectBlocking() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "two"
	}()

	// We use select to wait both of the values simultaneously
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// Timeouts are important for programs that connect to external resources or that otherwise need to bound execution time.
func selectTimeout() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	// awaits a value to be sent after the timeout of 1s; this case will run if the other operations take more than 1s.
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()
	select {
	// this case will be selected since it takes less than 3s
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}
}

// Basic sends and receives on channels are blocking.
// We can use `select` with default clause to implement non-blocking sends, receives and even non-blocking multi-way selects.
func selectNonBlocking() {
	messages := make(chan string)
	signals := make(chan bool)

	// If a value is available on `messages`, select will take the `<-messages` case with that value.
	// Otherwise, it will immediately take the `default` case.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// because the channel has no buffer and there is no receiver, the message cannot be sent.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// we can use multiple cases above the default clause to implement a multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

func SelectMain() {
	selectBlocking()

	selectTimeout()

	selectNonBlocking()
}
