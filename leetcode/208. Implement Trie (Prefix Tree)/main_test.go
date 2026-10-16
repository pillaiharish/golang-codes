package main

import "testing"

func TestTrie(t *testing.T) {
    trie:=Constructor()
    trie.Insert("apple")
    if !trie.Search("apple") {t.Fatal("full word missing")}
    if trie.Search("app") {t.Fatal("prefix must not be full word before insert")}
    if !trie.StartsWith("app") {t.Fatal("prefix missing")}
    trie.Insert("app")
    if !trie.Search("app") {t.Fatal("prefix did not become word")}
    trie.Insert("apple") // duplicate insertion
    if !trie.Search("apple") {t.Fatal("duplicate insertion broke word")}
    if trie.Search("apply") || trie.StartsWith("b") {t.Fatal("incorrect membership")}
    if !trie.StartsWith("") {t.Fatal("empty prefix should always match")}
}

func TestTrieEmptyWord(t *testing.T) {
    trie:=Constructor()
    if trie.Search("") {t.Fatal("empty word should not exist initially")}
    trie.Insert("")
    if !trie.Search("") {t.Fatal("empty word should be stored at root")}
}
