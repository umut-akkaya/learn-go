package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s string = "สวัสดี"
	const tayyar string = "tayyar"
	fmt.Println(len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x\n", s[i])
	}

	fmt.Println(utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	for idx, runeValue := range tayyar {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
}
