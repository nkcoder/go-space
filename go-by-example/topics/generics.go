package topics

import (
	"fmt"
	"iter"
	"slices"
)

func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}
	return -1
}

// List is a singly-linked list with values of any type
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push we can define methods on generic types just like we do on regular types, but we have to keep the type parameters in place
func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}
		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.tail.next
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

// All returns an iterator, which in Go is a function with a special signature
func (lst *List[T]) All() iter.Seq[T] {
	// the iterator function takes another function as a parameter, called yield by convention (the name can be arbitrary).
	// it will call yield for every element we want to iterate over, and note the return value for a potential early termination.
	return func(yield func(T) bool) {
		for e := lst.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Iteration doesn't require an underlying data structure, and doesn't even have to be finite.
// this funciton keeps running as long as yield keeps returning true.
func genFib() iter.Seq[int] {
	return func(yield func(int) bool) {
		a, b := 1, 1
		for {
			if !yield(a) {
				return
			}
			a, b = b, a+b
		}
	}
}

func Generics() {
	// when invoking generic functions, we can often rely on type inference
	s := []string{"foo", "bar", "baz"}
	fmt.Println("index of bar:", SliceIndex(s, "bar"))
	// you can also specify the type parameters explicitly, but it's not necessary
	_ = SliceIndex[[]string, string](s, "bar")

	l := List[int]{}
	l.Push(10)
	l.Push(13)
	l.Push(100)
	fmt.Println("all elements: ", l.GetAll())

	// we can use the iterator in a regular range loop
	lst2 := List[int]{}
	lst2.Push(11)
	lst2.Push(12)
	lst2.Push(30)
	for e := range lst2.All() {
		fmt.Println(e)
	}

	// packages like slices have a number of useful functions to work with iterators
	all := slices.Collect(lst2.All())
	fmt.Println(all)

	// once the loops hits break or an early return, the yield function passed to the iterator will return false
	for n := range genFib() {
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}
