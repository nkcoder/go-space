package main

import "fmt"

func Range() {
	nums := []int{3, 5, 6, 8, 9}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum of nums:", sum)

	// range on arrays and slices provides both the index and value for each entry
	for i, num := range nums {
		if num == 8 {
			fmt.Println("index of 8:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range can also iterate over just the keys of a map
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on strings iterates over Unicode code points
	// the first value is the starting byte index of the rune and the second is the rune itself
	for i, c := range "hello" {
		fmt.Println("index:", i, "char:", c)
	}

}
