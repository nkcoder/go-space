package main

import (
	"fmt"
	"time"
)

func Switch() {
	// basic switch
	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// use commas to separate multiple expressions in the same case statement
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// switch without an expression is an alternate way to express if/else logic
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	case t.Hour() < 18:
		fmt.Println("It's after noon")
	default:
		fmt.Println("It's night time")
	}

	// a type switch compares types instead of values
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
