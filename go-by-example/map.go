package main

import (
	"fmt"
	"maps"
)

func Map() {
	// Create an empty map
	m := make(map[string]int)

	// Set key/value pairs
	m["k1"] = 2
	m["k2"] = 8
	fmt.Println("map:", m, "len: ", len(m))

	// Get a value for a key
	// If the key does not exist, the zero value for the value type is returned
	v1 := m["k1"]
	v3 := m["k3"]
	fmt.Println("v1:", v1, "v3:", v3)

	// Delete a key/value pair
	delete(m, "k1")
	fmt.Println("map:", m, "len: ", len(m))

	// Remove all key/value pairs
	clear(m)
	fmt.Println("map:", m, "len: ", len(m))

	// The optional second return value indicates if the key was found
	// This can be used to disambiguate between missing keys and zero values
	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	// Declare and initialize a new map in the same line
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n, "len: ", len(n))

	// The `maps`` package contains a number of utility functions for working with maps
	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2) {
		fmt.Println("maps are equal")
	} else {
		fmt.Println("maps are not equal")
	}

	for k := range maps.Keys(n) {
		fmt.Println("key:", k)
	}
	for v := range maps.Values(n) {
		fmt.Println("value:", v)
	}

}
