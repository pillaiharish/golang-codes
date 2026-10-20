// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type Trie struct {
}

func Constructor() Trie {
	panic(todoMarker)
}

func (this *Trie) Insert(word string) {
	panic(todoMarker)
}

func (this *Trie) Search(word string) bool {
	panic(todoMarker)
}

func (this *Trie) StartsWith(prefix string) bool {
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
	obj.Insert("apple")
	fmt.Printf("search: %#v\n", obj.Search("apple"))
	fmt.Printf("search: %#v\n", obj.Search("app"))
	fmt.Printf("startsWith: %#v\n", obj.StartsWith("app"))
	obj.Insert("app")
	fmt.Printf("search: %#v\n", obj.Search("app"))
}

// Local fixture helpers only; do not submit to the judge.
