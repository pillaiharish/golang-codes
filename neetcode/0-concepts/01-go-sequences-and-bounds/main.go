package main

import (
	"fmt"
)

func main() {
	s := make([]int, 0, 3)
	s = append(s, 4, 8)
	view := s
	copyOfS := append([]int(nil), s...)
	s[0] = 9
	fmt.Println(len(s), cap(s), view, copyOfS)
}
