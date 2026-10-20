// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func exist(board [][]byte, word string) bool {
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
	board := [][]byte{[]byte{65, 66, 67, 69}, []byte{83, 70, 67, 83}, []byte{65, 68, 69, 69}}
	word := "ABCCED"
	fmt.Printf("result: %#v\n", exist(board, word))
}

// Local fixture helpers only; do not submit to the judge.
