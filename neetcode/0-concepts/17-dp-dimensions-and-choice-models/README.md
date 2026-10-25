# DP dimensions, choices, and counting conventions

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Two sequences naturally lead to prefix-pair states; interval tasks can use endpoint pairs; resource limits can become a remaining-budget dimension. The number of states times work per transition gives a first complexity estimate, but sparse states and output size may change it. A rolling array is valid only when overwritten entries are no longer needed.

Distinguish taking an item once from unlimited reuse, combinations from ordered sequences, and strict increases from nondecreasing order. These definitions change which transitions are legal and whether a state is counted more than once. Multiplicity is a contract decision, not a loop-order trick to memorize. Draw dependencies before optimizing storage.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
rows, cols := 3, 4
table := make([][]int, rows)
for i := range table { table[i] = make([]int, cols) }
table[1][2] = 7
fmt.Println(table[0][2], table[1][2], rows*cols)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **0 7 12**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Draw dependencies for an abstract two-dimensional table. Identify which old cells must remain available before reusing a row.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Reusing row backing arrays; switching combination and permutation counting; compressing dimensions before proving dependency safety; incorrect empty-prefix initialization.

## Questions that use this building block

- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)
- [#416 Partition Equal Subset Sum](../../416.%20Partition%20Equal%20Subset%20Sum/README.md)
- [#62 Unique Paths](../../62.%20Unique%20Paths/README.md)
- [#1143 Longest Common Subsequence](../../1143.%20Longest%20Common%20Subsequence/README.md)
- [#309 Best Time to Buy and Sell Stock with Cooldown](../../309.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock%20with%20Cooldown/README.md)
- [#518 Coin Change II](../../518.%20Coin%20Change%20II/README.md)
- [#494 Target Sum](../../494.%20Target%20Sum/README.md)
- [#97 Interleaving String](../../97.%20Interleaving%20String/README.md)
- [#329 Longest Increasing Path in a Matrix](../../329.%20Longest%20Increasing%20Path%20in%20a%20Matrix/README.md)
- [#115 Distinct Subsequences](../../115.%20Distinct%20Subsequences/README.md)
- [#72 Edit Distance](../../72.%20Edit%20Distance/README.md)
- [#312 Burst Balloons](../../312.%20Burst%20Balloons/README.md)
- [#10 Regular Expression Matching](../../10.%20Regular%20Expression%20Matching/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
