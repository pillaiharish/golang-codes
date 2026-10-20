package main

import (
	"fmt"
)

func main() {
	x := []int{4, 6, 8, 10}
	left, right := 1, 3
	fmt.Println(x[left:right], right-left)
	left++
	fmt.Println(x[left:right], right-left)
}
