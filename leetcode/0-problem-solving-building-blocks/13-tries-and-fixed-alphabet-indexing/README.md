# 13 — Trie nodes and character index tables

[← Foundations index](../README.md) · [← Problem index](../../README.md)

## Recognize the building block
A trie maps each prefix to a path. Fixed child-pointer arrays provide fast indexing for small alphabets; `isEnd` separates whole words from prefixes.

## Operations you must be able to write
- Use `[26]*Node` for lowercase ASCII letters, and calculate index `word[i]-'a'`.
- On Insert, create a missing child and move to that child; after the loop mark `isEnd=true`.
- On Search, return false for a missing child and otherwise return `isEnd` at the last node.
- On StartsWith, a traversable prefix is enough; don't require `isEnd`.
- For broader alphabets, use `map[rune]*Node` and loop over runes, with different time/memory tradeoffs.

## Complete small Go example
```go
package main
import "fmt"
type Node struct { Next [26]*Node; End bool }
func main(){
    root:=&Node{}
    n:=root
    for _,b:=range []byte("app"){
        i:=b-'a'
        if n.Next[i]==nil {n.Next[i]=&Node{}}
        n=n.Next[i]
    }
    n.End=true
    fmt.Println(root.Next[0].Next['p'-'a'].Next['p'-'a'].End) // true
}
```

## Dry run
Insert `app`: root→a→p→p then mark final p as End. The prefix `ap` exists but is not itself a full word.

## Tiny practice task
Insert `apple`, then answer `Search(app)=false`, `StartsWith(app)=true`. After Insert(app), both become true.

## Common mistakes
- Indexing with non-a–z bytes can panic; never claim this implementation supports all Unicode.
- A node's existence does not prove a complete word was inserted.
- A `[26]*Node` array uses many child pointers even when most links are empty.

## Problem links from this repository
- [#208 Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go)
- [#242 Valid Anagram](../../242.%20Valid%20Anagram/main.go)
- [#49 Group Anagrams](../../49.%20Group%20Anagrams/main.go)

## Blank-screen check
Close this note. Restate the invariant, retype the operations, then solve the first linked problem without looking.
