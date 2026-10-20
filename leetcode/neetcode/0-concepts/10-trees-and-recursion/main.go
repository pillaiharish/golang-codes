package main

import (
	"fmt"
)

func main() {
	type Branch struct {
		Value       int
		Left, Right *Branch
	}
	root := &Branch{Value: 8, Left: &Branch{Value: 3}}
	fmt.Println(root.Value, root.Left.Value, root.Right == nil)
}
