package topics

import (
	"fmt"
	"unicode/utf8"
)

func String() {
	// in Go, the concept a character is called a rune - it is an integer type that represents a Unicode code point
	const s = "你好吗"
	//strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	fmt.Println("len: ", len(s)) // 9

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i]) // e4bda0e5a5bde59097
	}

	fmt.Println()

	fmt.Println("Rune count: ", utf8.RuneCountInString(s)) // 3

	// a range loop handles strings especially and decodes each rune along with its offset in the string
	for idx, runeValue := range s {
		// U+4F60 '你' starts at 0
		// U+597D '好' starts at 3
		// U+5417 '吗' starts at 6
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// we can achieve the same iteration by using utf8.DecodeRuneInString
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		// U+4F60 '你' starts at 3
		// U+597D '好' starts at 3
		// U+5417 '吗' starts at 3
		fmt.Printf("%#U starts at %d\n", runeValue, width)
		w = width
	}
}
