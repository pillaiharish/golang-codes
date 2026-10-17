package main

import "fmt"

func main() {
	var n int
	var s []int
	var m map[string]int
	var ch chan int
	fmt.Println(n, s == nil, m == nil, ch == nil) // 0 true true true
	s = append(s, 7)
	v, ok := m["go"]
	fmt.Println(s, v, ok) // [7] 0 false
	m = make(map[string]int)
	m["go"]++
	ch = make(chan int, 3)
	fmt.Println(m["go"], len(ch), cap(ch)) // 1 0 3
	var p *int
	var x any = p
	fmt.Println(p == nil, x == nil) // true false
}
