// Unsolved starter. Judge interface verified from LeetCode.
package main

import (
	"fmt"
)

// BEGIN JUDGE CODE: implement these TODO bodies.
type Twitter struct {
}

func Constructor() Twitter {
	panic(todoMarker)
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	panic(todoMarker)
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	panic(todoMarker)
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	panic(todoMarker)
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
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
	obj.PostTweet(1, 5)
	fmt.Printf("getNewsFeed: %#v\n", obj.GetNewsFeed(1))
	obj.Follow(1, 2)
	obj.PostTweet(2, 6)
	fmt.Printf("getNewsFeed: %#v\n", obj.GetNewsFeed(1))
	obj.Unfollow(1, 2)
	fmt.Printf("getNewsFeed: %#v\n", obj.GetNewsFeed(1))
}

// Local fixture helpers only; do not submit to the judge.
