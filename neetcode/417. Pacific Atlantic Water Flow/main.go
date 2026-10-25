// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func pacificAtlantic(heights [][]int) [][]int {
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
	heights := [][]int{[]int{1, 2, 2, 3, 5}, []int{3, 2, 3, 4, 4}, []int{2, 4, 5, 3, 1}, []int{6, 7, 1, 4, 5}, []int{5, 1, 1, 2, 4}}
	fmt.Printf("result: %#v\n", pacificAtlantic(heights))
}

// Local fixture helpers only; do not submit to the judge.
