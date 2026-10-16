# 208 — Implement Trie (Prefix Tree)

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [04 strings bytes runes and parsing](../0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md)
2. [11 structs pointers methods and linked lists](../0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md)
3. [13 tries and fixed alphabet indexing](../0-problem-solving-building-blocks/13-tries-and-fixed-alphabet-indexing/README.md)

[Full mapping for all 23 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-23-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: trie paths](../0-problem-solving-building-blocks/README.md#6-tries-are-character-by-character-paths) · [Go solution](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/implement-trie-prefix-tree/)

## Recognition

**Clue:** repeated `Insert(word)`, `Search(exactWord)` and `StartsWith(prefix)`. A hash set does exact membership but a trie makes prefix traversal natural.

**Mental trigger:** **One character = one edge; complete word requires isEnd.**

## Intuition

Words with the same prefix share a path from the root. A node contains 26 child pointers and an `isEnd` flag. Traversing an existing prefix does not mean a complete word was inserted.

Picture after inserting `apple`:

```text
root → a → p → p → l → e*     (* means isEnd=true)
```

`Search("app")` = false; `StartsWith("app")` = true. After inserting `app`, the node for the second p gains `isEnd=true`.

## Invariant

After processing the first `i` bytes of a valid lowercase word, `current` points to the trie node representing **exactly that prefix**. A complete word matches only if the final node has `isEnd=true`.

## Recipe

1. Constructor allocates a root node.
2. **Insert:** start at root; for every character compute `index := word[i]-'a'`; create a child if missing; advance; finally mark `isEnd`.
3. **Search:** follow children; missing edge → false; after final character return `isEnd`.
4. **StartsWith:** same traversal, but if all edges exist return true regardless of `isEnd`.

### Dry run

| Operation | Path | End flag | Answer |
|---|---|---|---|
| Insert("apple") | a→p→p→l→e | mark e | — |
| Search("apple") | a→p→p→l→e | true | true |
| Search("app") | a→p→p | false | false |
| StartsWith("app") | a→p→p | irrelevant | true |
| Insert("app") | existing a→p→p | mark second p | — |
| Search("app") | a→p→p | true | true |

## Complexity and contract

If a query has L characters: Insert/Search/StartsWith O(L) time. The trie stores up to O(total inserted characters) nodes, but an array of 26 pointers per node means **constant but potentially large per-node overhead**.

**Important:** These implementations assume **lowercase ASCII a–z only**, not arbitrary UTF-8. Any other byte can cause out-of-range indexing; a production API should validate or use `map[rune]*node`. Empty prefix always exists; if `Insert("")` is called, root's `isEnd` makes `Search("")` true.

## Alternatives

- Map set for exact word membership only.
- Sorted string slice + binary search and range lookup for mostly immutable dictionaries.
- Rune/map-based trie for broader alphabets, with more allocation and lookup overhead.

## 30-second spoken answer

> A trie stores common prefixes as shared paths. Each node has pointers for possible next letters and an isEnd flag. Insertion creates missing nodes and marks the terminal node. Exact Search requires both the path and isEnd; StartsWith only requires the path. All operations cost O(L), independent of how many words already exist.

## Cold recall

1. Why is the `isEnd` flag necessary?
2. How does `Search("app")` differ from `StartsWith("app")`?
3. Why is `word[i]-'a'` only safe for a narrow input contract?
4. Trace duplicate and empty string insertion.

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
