package topics

import (
	"fmt"
	"slices"
)

func Slice() {
	var s []string
	fmt.Println("uninitialized slice:", s, s == nil, len(s), cap(s)) // [] true 0 0

	s = make([]string, 3)
	fmt.Println("initialized slice:", s, s == nil, len(s), cap(s)) // [ ] false 3 3

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s, s[2], len(s), cap(s)) // [a b c] c 3 3

	s = append(s, "d")
	fmt.Println(s, len(s), cap(s)) // [a b c d] 4 6
	s = append(s, "e", "f", "g")
	fmt.Println(s, len(s), cap(s)) // [a b c d e f g] 7 12

	// slices can also be copied
	r := make([]string, len(s))
	copy(r, s)
	fmt.Println("after copy, r: ", r, len(r), cap(r)) // [a b c d e f g] 7 7

	l1 := s[2:5]
	fmt.Println("l1: ", l1) // [c d e]

	l2 := s[:5]
	fmt.Println("l2: ", l2) // [a b c d e]

	l3 := s[2:]
	fmt.Println("l3: ", l3) // [c d e f g]

	// there are many utility functions in the slices package
	t1 := []string{"a", "b", "c"}
	t2 := []string{"d", "e", "f"}
	if slices.Equal(t1, t2) {
		fmt.Println("t1 and t2 are equal")
	} else {
		fmt.Println("t1 and t2 are not equal") // not equal
	}

	// for multidimensional slices, the length of each inner slice can vary
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := 2 * (i + 1)
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}

	// test append
	scores := []int{1, 2, 3}
	appendScores(scores, 4, 5, 6)
	fmt.Println(`scores after appendScores: `, scores) // [1 2 3], scores is unchanged
}

// functions operates on copies of arguments
func appendScores(s []int, values ...int) {
	fmt.Println("before append", s) // [1 2 3]
	s = append(s, values...)        // it depends if the underlying array has sufficient capacity
	fmt.Println("after append", s)  // [1 2 3 4 5 6]
}
