// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func longestIncreasingPath(matrix [][]int) int {
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
	matrix := [][]int{[]int{9, 9, 4}, []int{6, 6, 8}, []int{2, 1, 1}}
	fmt.Printf("result: %#v\n", longestIncreasingPath(matrix))
}

// Local fixture helpers only; do not submit to the judge.
