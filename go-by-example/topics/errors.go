// Package topics go topics in Go by example
package topics

import (
	"errors"
	"fmt"
)

// by convention, errors are the last return value and have type error, a built-in interface
func f(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message
		return -1, errors.New("cannot work with 42")
	}

	// a nil value in the error position indicates that there was no error
	return arg + 3, nil
}

// a sentinel error is a predeclared variable that is used to signify a specific error condition
var (
	ErrOutOfTea = fmt.Errorf("no more team available")
	ErrPower    = fmt.Errorf("cannot boil water")
)

func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		// we can wrap errors with higher-level errors to add context, %w (similar to %v) is primarily used with `fmt.Errorf` for error wrapping to create a error chain
		// %w causes the returned error to implement an `Unwrap` method to return the wrapped error.
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

// it is possible to use custom types as errors by implmenting the Error() method on them
// a custom error type usually has a suffix 'Error'
type argError struct {
	arg     int
	message string
}

// adding the Error method makes argError implement the error interface
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f2(arg int) (int, error) {
	if arg == 50 {
		return -1, &argError{arg, "cannot work with it"}
	}
	return arg + 3, nil
}

func ErrorMain() {
	for _, i := range []int{7, 42} {
		// it is common to use an inline error check in the if line
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// errors.Is checks that a given error (or any error in its chain) matches a specific error value.
	// this is especially useful with wrapped or nested errors
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}

	_, err := f2(50)
	var ae *argError
	// errors.As is a more advanced version of errors.Is: it checks that a given error (or any error in its chain) matches a specific error type and converts to a value of the type, returning true.
	// if there is no match, it returns false
	if errors.As(err, &ae) {
		fmt.Println(ae.arg, ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
