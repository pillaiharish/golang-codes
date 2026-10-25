// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
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
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Printf("result: %#v\n", findMedianSortedArrays(nums1, nums2))
}

// Local fixture helpers only; do not submit to the judge.
