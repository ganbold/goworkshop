package main

import (
	"fmt"
	"unicode/utf8"
)

// in golan char == rune
// https://go.dev/blog/strings
func main_14() {
	const s = "өнөөдрийн үр дүн"

	fmt.Println("Len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	switch r {
	case 'н':
		fmt.Println("н үсэгтэй тааралдлаа")
	case 'д':
		fmt.Println("д үсэгтэй тааралдлаа")
	}
}
