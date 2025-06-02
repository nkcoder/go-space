package main

import (
	"fmt"
)

// No additional imports needed since Hello and Values
// are in the same package (main)

func main() {
	Hello()
	printLine("hello")

	Values()
	printLine("values")

	Variables()
	printLine("variables")

	Constants()
	printLine("constants")

	ForLoop()
	printLine("for loop")

	IfElse()
	printLine("if else")

	Switch()
	printLine("switch")

	Arrays()
	printLine("arrays")

	Map()
	printLine("map")

	Function()
	printLine("function")

	Range()
	printLine("range")

	Pointer()
	printLine("pointer")

	String()
	printLine("string")

	Struct()
	printLine("struct")

	Interface()
	printLine("interface")

	Enum()
	printLine("enum")

	StructEmbedding()
	printLine("struct embedding")
}

func printLine(name string) {
	fmt.Printf("------------%s------------------\n\n", name)
}
