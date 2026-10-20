// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func maxAreaOfIsland(grid [][]int) int {
	panic(todoMarker)
}

// END JUDGE CODE

// Local harness: catches only the intentional unimplemented marker.
// Unexpected panics are rethrown. A TODO run is not a correctness test.
type localTODO struct{}

var todoMarker = &localTODO{}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if r == todoMarker {
				fmt.Println("TODO: implement the judge function/methods, then run again.")
				return
			}
			panic(r)
		}
	}()
	localDemo()
}
func localDemo() {
	grid := [][]int{[]int{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, []int{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, []int{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, []int{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	fmt.Printf("result: %#v\n", maxAreaOfIsland(grid))
}

// Local fixture helpers only; do not submit to the judge.
