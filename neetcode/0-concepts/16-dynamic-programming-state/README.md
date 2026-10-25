# Dynamic programming state, base cases, and evaluation order

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Dynamic programming reuses results for overlapping subproblems. A state must include every fact that can affect future choices. Define the meaning of dp[i] or dp[i][j] in a sentence before writing a recurrence. Memoization computes states on demand; tabulation evaluates them in a chosen order. Both need termination and correct base cases.

Zero may mean a valid result, unreachable state, or uncomputed state; these meanings cannot safely share a representation without justification. Boolean feasibility, minimum cost, maximum reward, and number of ways use different combination rules. An empty prefix often needs a meaningful base case rather than being an error. This lesson teaches storage and state reasoning; it deliberately supplies no problem-specific recurrence.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type State struct { I, J int }
memo := make(map[State]int)
k := State{I: 2, J: 1}
v, exists := memo[k]
fmt.Println(v, exists)
memo[k] = 0
v, exists = memo[k]
fmt.Println(v, exists)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **0 false; 0 true**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Define a state for processing two prefixes, list its possible dimensions, and mark which transitions would be acyclic.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Using zero as both computed result and missing marker; missing state dimensions; computing before dependencies; claiming O(n) for an n-by-m state table.

## Questions that use this building block

- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)
- [#70 Climbing Stairs](../../70.%20Climbing%20Stairs/README.md)
- [#746 Min Cost Climbing Stairs](../../746.%20Min%20Cost%20Climbing%20Stairs/README.md)
- [#198 House Robber](../../198.%20House%20Robber/README.md)
- [#213 House Robber II](../../213.%20House%20Robber%20II/README.md)
- [#5 Longest Palindromic Substring](../../5.%20Longest%20Palindromic%20Substring/README.md)
- [#647 Palindromic Substrings](../../647.%20Palindromic%20Substrings/README.md)
- [#91 Decode Ways](../../91.%20Decode%20Ways/README.md)
- [#322 Coin Change](../../322.%20Coin%20Change/README.md)
- [#152 Maximum Product Subarray](../../152.%20Maximum%20Product%20Subarray/README.md)
- [#139 Word Break](../../139.%20Word%20Break/README.md)
- [#300 Longest Increasing Subsequence](../../300.%20Longest%20Increasing%20Subsequence/README.md)
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
