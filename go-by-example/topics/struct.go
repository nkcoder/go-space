package topics

import "fmt"

type person struct {
	name string
	age  int
}

// Go is a garbage collected language, you can safely return a pointer to a local variable
// it will only be cleaned up by the garbage collector when there are no active references to it.
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

// Go supports methods defined on struct types
type rect struct {
	width, height int
}

// the method has a receiver type of *rect
func (r *rect) area() int {
	return r.width * r.height
}

// methods can be defined for either pointer or value receiver types
func (r rect) perimeter() int {
	return 2*r.width + 2*r.height
}

func Struct() {

	// create a new struct
	fmt.Println(person{"Bob", 20})

	// you can name the fields when initializing a struct
	fmt.Println(person{name: "Alice", age: 30})

	// omitted fields will be zero-valued
	fmt.Println(person{name: "Bob"})

	// create a pointer to the struct
	fmt.Println(&person{name: "Daniel", age: 60})

	// it's idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(newPerson("John"))

	p1 := person{name: "Anthony", age: 20}
	fmt.Println(p1.name, p1.age)

	pp1 := &p1
	fmt.Println(pp1.age, pp1.name)

	pp1.age = 70
	fmt.Println(pp1.age)

	// anonymous struct type
	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)

	/*
		Go automatically handles conversions between values and pointers for method calls.
		You many want to use a pointer receiver type to avoid copying on method calls or to allow the method to mutate the receiving struct
	*/
	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perimeter: ", r.perimeter())

	rp := r
	fmt.Println("area: ", rp.area())
	fmt.Println("perimeter: ", rp.perimeter())
}
