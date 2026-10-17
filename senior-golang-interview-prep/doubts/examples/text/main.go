package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "Aé界"
	fmt.Println(len(s), utf8.RuneCountInString(s)) // 6 3
	fmt.Printf("% x\n", []byte(s))
	for offset, r := range s {
		fmt.Printf("%d %U %c\n", offset, r, r)
	}
	invalid := string([]byte{0xff, 'A'})
	r, width := utf8.DecodeRuneInString(invalid)
	fmt.Printf("valid=%t first=%U width=%d\n", utf8.ValidString(invalid), r, width)
}
