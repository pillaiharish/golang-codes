package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		} else {
			return intervals[i][0] > intervals[j][0]
		}
	})
	result := 0
	p1, p2 := -1, -1
	fmt.Println(intervals)
	for _, intv := range intervals {
		if intv[0] > p2 {
			result += 2
			p2 = intv[1]
			p1 = p2 - 1
		} else if intv[0] > p1 {
			result += 1
			p1 = p2
			p2 = intv[1]
		}
	}
	return result
}

func main() {
	intervals := [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}
	fmt.Println(intersectionSizeTwo(intervals))
}

// Output:
// harish $ go run main.go
// [[1 3] [1 4] [3 5] [2 5]]
// 3
