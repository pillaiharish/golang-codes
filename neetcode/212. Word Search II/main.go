// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func findWords(board [][]byte, words []string) []string {
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
	board := [][]byte{[]byte{111, 97, 97, 110}, []byte{101, 116, 97, 101}, []byte{105, 104, 107, 114}, []byte{105, 102, 108, 118}}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Printf("result: %#v\n", findWords(board, words))
}

// Local fixture helpers only; do not submit to the judge.
