package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
	r, c := 1, 2
	fmt.Println(len(matrix), len(matrix[0]), matrix[r][c])
	matrix[0][0], matrix[1][0] = matrix[1][0], matrix[0][0]
	fmt.Println(matrix)
}
