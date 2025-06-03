package main

import (
	"fmt"
)

// No additional imports needed since Hello and Values
// are in the same package (main)

func main() {
	Slice()
}

func printLine(name string) {
	fmt.Printf("------------%s------------------\n\n", name)
}
