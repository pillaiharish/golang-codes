package main

import (
	"fmt"
)

func main() {
	parent := make([]int, 4)
	size := make([]int, 4)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	fmt.Println(parent, size)
	indegree := make([]int, 3)
	indegree[2]++
	fmt.Println(indegree)
}
