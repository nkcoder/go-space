package topics

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	const s = "你好吗"
	//strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	fmt.Println("len: ", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	// a range loop handles strings especially and decodes each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// we can achieve the same iteration by using utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, width)
		w = width
	}
}
