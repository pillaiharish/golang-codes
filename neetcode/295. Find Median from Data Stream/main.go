// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type MedianFinder struct {
}

func Constructor() MedianFinder {
	panic(todoMarker)
}

func (this *MedianFinder) AddNum(num int) {
	panic(todoMarker)
}

func (this *MedianFinder) FindMedian() float64 {
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
	obj := Constructor()
	obj.AddNum(1)
	obj.AddNum(2)
	fmt.Printf("findMedian: %#v\n", obj.FindMedian())
	obj.AddNum(3)
	fmt.Printf("findMedian: %#v\n", obj.FindMedian())
}

// Local fixture helpers only; do not submit to the judge.
