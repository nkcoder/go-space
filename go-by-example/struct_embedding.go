package main

import "fmt"

// struct embedding
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num: %d", b.num)
}

type container struct {
	base
	str string
}

func StructEmbedding() {
	// struct embedding
	con := container{
		base: base{num: 10},
		str:  "i'm a container",
	}
	// we can access the base's fields directly on container
	fmt.Printf("con={num: %v, str: %v}\n", con.num, con.str)
	// we can also use the full path to access the base's fields
	fmt.Println("con.base.num: ", con.base.num)
	// the methods of base also become methods of container
	fmt.Println("con.describe: ", con.describe())

	type describer interface {
		describe() string
	}
	var d describer = con
	fmt.Println("d.describe: ", d.describe())
}
