package main

import "fmt"

var saved *int

//go:noinline
func local() int {
	p := new(int)
	*p = 7
	return *p
}

//go:noinline
func retain() {
	n := 7
	saved = &n
}

func main() {
	retain()
	fmt.Println(local(), *saved) // 7 7
}
