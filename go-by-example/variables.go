package main

import "fmt"

func Variables() {
	// var declares 1 or more variables
	var a = "initial"
	fmt.Println(a)

	// multiple variables can be declared at once
	var b, c int = 1, 2
	fmt.Println(b, c)

	// type is inferred
	var d = true
	fmt.Println(d)

	// declared without initialization is zero-valued
	var e int
	fmt.Println(e)

	// can only be used inside functions
	f := "apple"
	fmt.Println(f)
}
