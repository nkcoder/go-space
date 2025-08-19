package topics

import "fmt"

func Arrays() {
	var a [5]int
	fmt.Println("Empty array:", a)

	a[0] = 1
	a[3] = 4
	fmt.Println("a:", a)
	fmt.Println("a length:", len(a))
	fmt.Println("a capacity:", cap(a))
	fmt.Println("a first element:", a[0])

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b:", b)

	// you can also use ... to let the compiler count the elements
	c := [...]int{2, 4, 6, 8}
	fmt.Println("c:", c)

	// if you specify the index with :, the elements in between will be zeroed
	d := [...]int{100, 4: 400, 500}
	fmt.Println("d:", d)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:", twoD)
}
