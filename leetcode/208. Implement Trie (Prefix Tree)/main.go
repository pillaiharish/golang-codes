package main

import "fmt"

// A trie stores prefixes along paths. isEnd marks complete words.
// Contract: only lowercase ASCII letters 'a' through 'z'.
type node struct {
    children [26]*node
    isEnd bool
}

type Trie struct { root *node }

func Constructor() Trie { return Trie{root:&node{}} }

func (t *Trie) Insert(word string) {
    current := t.root
    for i:=0;i<len(word);i++ {
        index := word[i]-'a'
        if current.children[index]==nil { current.children[index]=&node{} }
        current = current.children[index]
    }
    current.isEnd = true
}

func (t *Trie) Search(word string) bool {
    current := t.root
    for i:=0;i<len(word);i++ {
        index := word[i]-'a'
        if current.children[index]==nil { return false }
        current = current.children[index]
    }
    return current.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
    current := t.root
    for i:=0;i<len(prefix);i++ {
        index := prefix[i]-'a'
        if current.children[index]==nil { return false }
        current = current.children[index]
    }
    return true
}

func main() {
    trie := Constructor()
    trie.Insert("apple")
    fmt.Println(trie.Search("apple"),trie.Search("app"),trie.StartsWith("app")) // true false true
}
