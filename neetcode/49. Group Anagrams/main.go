// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func groupAnagrams(strs []string) [][]string {
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
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Printf("result: %#v\n", groupAnagrams(strs))
}

// Local fixture helpers only; do not submit to the judge.
