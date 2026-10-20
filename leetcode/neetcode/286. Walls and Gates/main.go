// Unsolved starter. Judge interface adapted locally; premium Go starter unavailable.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func wallsAndGates(rooms [][]int) {
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
	rooms := [][]int{[]int{2147483647, -1, 0, 2147483647}, []int{2147483647, 2147483647, 2147483647, -1}, []int{2147483647, -1, 2147483647, -1}, []int{0, -1, 2147483647, 2147483647}}
	wallsAndGates(rooms)
	fmt.Printf("mutated input: %#v\n", rooms)
}

// Local fixture helpers only; do not submit to the judge.
