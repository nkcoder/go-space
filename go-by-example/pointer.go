package main

import "fmt"

// pass by value, so the function will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 2
}

// pass by reference: it will dereference the pointer from its memory address to the current value at that address
func zeroptr(iptr *int) {
	*iptr = 3
}

func Pointer() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr: ", i)

	// pointers can be printed too
	fmt.Println("pointer: ", &i)

}
