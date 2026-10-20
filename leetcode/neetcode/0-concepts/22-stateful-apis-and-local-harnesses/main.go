package main

import (
	"fmt"
)

func main() {
	type Counter struct{ Value int }
	c := &Counter{}
	c.Value++
	c.Value++
	fmt.Println(c.Value)
	other := &Counter{}
	fmt.Println(other.Value)
}
