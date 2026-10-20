package main

import (
	"fmt"
)

func main() {
	path := []int{4}
	path = append(path, 7)
	snapshot := append([]int(nil), path...)
	path = path[:len(path)-1]
	path = append(path, 9)
	fmt.Println(path, snapshot)
}
