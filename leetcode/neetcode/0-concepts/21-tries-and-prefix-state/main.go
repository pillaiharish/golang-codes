package main

import (
	"fmt"
)

func main() {
	type PrefixNode struct {
		Terminal bool
		Children map[byte]*PrefixNode
	}
	root := &PrefixNode{Children: make(map[byte]*PrefixNode)}
	root.Children['a'] = &PrefixNode{Terminal: true}
	fmt.Println(root.Terminal, root.Children['a'].Terminal)
}
