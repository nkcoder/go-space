package topics

import "fmt"

// pass by value, so the function will get a copy of ival distinct from the one in the calling function
func zeroval(ival int) {
	ival = 2
}

// In go, it's always pass by value
// We still have a copy of the argument, but it points to the same memory address
// so assigning to the dereferenced pointer will change the value at the referenced address
func zeroptr(iptr *int) {
	*iptr = 3
}

func Pointer() {
	i := 1
	fmt.Println("initial: ", i)

	zeroval(i)
	fmt.Println("zeroval: ", i) // 1

	zeroptr(&i)
	fmt.Println("zeroptr: ", i) // 3

	// pointers can be printed too
	fmt.Println("pointer: ", &i) // 0xc0000140a0

}
