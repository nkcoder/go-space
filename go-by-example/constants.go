package main

import (
	"fmt"
	"math"
)

// const declares a constant value, can appear anywhere a var can
const s string = "constant"

func Constants() {
	fmt.Println(s)

	const n = 500000000

	// a numeric constant has no type until it's give none, such as by an explicit conversion
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// a number can be given a type by using it in a context that requires one, such as a variable assignment or function call
	fmt.Println(math.Sin(n))
}
