package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := "Aé"
	fmt.Println(len(s), utf8.RuneCountInString(s))
	for i, r := range s {
		fmt.Printf("%d %U\n", i, r)
	}
	fmt.Println(strconv.Itoa(65))
}
