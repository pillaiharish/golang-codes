package main

import (
	"fmt"
)

func main() {
	stack := []int{2, 7}
	stack = append(stack, 9)
	last := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	queue := []int{3, 5, 8}
	head := 0
	first := queue[head]
	head++
	fmt.Println(last, stack, first, queue[head:])
}
