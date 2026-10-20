# Trie nodes, prefixes, and terminal markers

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A trie represents sequences by paths through labeled edges. A node can mark a complete word and also have children for longer words; a reachable prefix is not automatically a stored word. Fixed child arrays fit a known small alphabet, while maps support a wider dynamic alphabet at different memory costs.

A wildcard search may have to consider several outgoing edges, but the wildcard's meaning comes from the contract: one character is not the same as any-length suffix. Node allocation and nil checks are ordinary pointer operations. Learn the representation here; implement insertion/search only when solving the design questions yourself.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type PrefixNode struct { Terminal bool; Children map[byte]*PrefixNode }
root := &PrefixNode{Children: make(map[byte]*PrefixNode)}
root.Children['a'] = &PrefixNode{Terminal: true}
fmt.Println(root.Terminal, root.Children['a'].Terminal)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **false true**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw paths for a short word and a longer word sharing its prefix. Mark exactly which nodes are terminal.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Equating prefix existence with word membership; nil-map writes; using a 26-slot array outside lowercase ASCII; treating a dot wildcard as an arbitrary-length match.

## Questions that use this building block

- [#208 Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md)
- [#211 Design Add and Search Words Data Structure](../../211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/README.md)
- [#212 Word Search II](../../212.%20Word%20Search%20II/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/13-tries-and-fixed-alphabet-indexing/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
