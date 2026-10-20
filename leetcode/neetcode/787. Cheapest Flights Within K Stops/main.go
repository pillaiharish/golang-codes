// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
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
	n := 4
	flights := [][]int{[]int{0, 1, 100}, []int{1, 2, 100}, []int{2, 0, 100}, []int{1, 3, 600}, []int{2, 3, 200}}
	src := 0
	dst := 3
	k := 1
	fmt.Printf("result: %#v\n", findCheapestPrice(n, flights, src, dst, k))
}

// Local fixture helpers only; do not submit to the judge.
