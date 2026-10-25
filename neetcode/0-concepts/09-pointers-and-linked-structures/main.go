package main

import (
	"fmt"
)

func main() {
	type Cell struct {
		Value int
		Next  *Cell
	}
	a := &Cell{Value: 4}
	b := &Cell{Value: 4}
	a.Next = b
	fmt.Println(a == b, a.Value == b.Value, a.Next == b)
}
