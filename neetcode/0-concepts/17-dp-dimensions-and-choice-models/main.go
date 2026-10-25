package main

import (
	"fmt"
)

func main() {
	rows, cols := 3, 4
	table := make([][]int, rows)
	for i := range table {
		table[i] = make([]int, cols)
	}
	table[1][2] = 7
	fmt.Println(table[0][2], table[1][2], rows*cols)
}
