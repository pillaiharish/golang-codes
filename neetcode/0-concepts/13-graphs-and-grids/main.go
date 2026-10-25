package main

import (
	"fmt"
)

func main() {
	adj := make([][]int, 3)
	adj[0] = append(adj[0], 2)
	adj[2] = append(adj[2], 0)
	fmt.Println(adj[0], adj[1], adj[2])
	cell := [2]int{1, 2}
	seen := map[[2]int]bool{cell: true}
	fmt.Println(seen[cell])
}
