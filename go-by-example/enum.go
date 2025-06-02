package main

import "fmt"

type ServerState int

/*
*
We use the `iota` keyword, which is a special Go construct for creating enumerated constants.

- iota is a counter that starts at 0 and increments by 1 for each constant in the block.
It resets to 0 when a new const block is started.

- StateIdle, StateConnected, StateError, and StateRetrying are constants of type ServerState because the first constant (StateIdle) explicitly declares its type as ServerState.
All subsequent constants in the block inherit the same type (ServerState).
Incremental Values:

- StateIdle is assigned the value 0 (the initial value of iota). StateConnected is assigned the value 1 (the next value of iota). Similarly, StateError gets 2, and StateRetrying gets 3.
*/
const (
	StateIdle ServerState = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdle:      "Idle",
	StateConnected: "Connected",
	StateError:     "Error",
	StateRetrying:  "Retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}

func Enum() {

	// If we have a type of int, we cannot pass it to transition - the compiler will complain about type mismatch
	ns := transition(StateIdle)
	fmt.Println(ns)

	next := transition(ns)
	fmt.Println(next)
}
