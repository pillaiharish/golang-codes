package main

import (
	"fmt"
)

func main() {
	type Edge struct{ To, Cost int }
	adj := [][]Edge{{{To: 1, Cost: 6}}, {}}
	routeCost := 4
	candidate := routeCost + adj[0][0].Cost
	fmt.Println(adj[0][0].To, candidate)
}
