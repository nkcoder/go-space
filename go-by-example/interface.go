package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

// to implement an interface in Go, we just need to implement all the methods in the interface
func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, then we can call methods that are in the named interface
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}

func detectCircle(g geometry) {
	// type assertion: asserts that the interface holds the concrete type and assigns the underlying value to the variable.
	// the boolean value reports whether the assertion succeeded
	if c, ok := g.(circle); ok {
		fmt.Println("circle with radius: ", c.radius)
	}
}

func Interface() {
	r := rectangle{width: 10, height: 5}
	c := circle{radius: 5}

	measure(r)
	measure(c)

	detectCircle(r)
	detectCircle(c)
}
