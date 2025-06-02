package main

import "fmt"

func Values() {
	// string
	fmt.Println("go" + "lang")

	// int and float
	fmt.Println("1 + 1 = ", 1+1)
	fmt.Println("7.0 / 3.0 = ", 7.0/3.0)

	// boolean
	fmt.Println(true && false)
	fmt.Println(!true)
	fmt.Println(true || false)
}
