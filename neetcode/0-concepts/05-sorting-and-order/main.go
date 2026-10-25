package main

import (
	"fmt"
	"sort"
)

func main() {
	x := []int{8, 3, 5}
	y := append([]int(nil), x...)
	sort.Ints(y)
	fmt.Println(x, y)
}
