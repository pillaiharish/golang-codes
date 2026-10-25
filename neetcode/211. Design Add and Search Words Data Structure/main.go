// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type WordDictionary struct {
}

func Constructor() WordDictionary {
	panic(todoMarker)
}

func (this *WordDictionary) AddWord(word string) {
	panic(todoMarker)
}

func (this *WordDictionary) Search(word string) bool {
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
	obj.AddWord("bad")
	obj.AddWord("dad")
	obj.AddWord("mad")
	fmt.Printf("search: %#v\n", obj.Search("pad"))
	fmt.Printf("search: %#v\n", obj.Search("bad"))
	fmt.Printf("search: %#v\n", obj.Search(".ad"))
	fmt.Printf("search: %#v\n", obj.Search("b.."))
}

// Local fixture helpers only; do not submit to the judge.
