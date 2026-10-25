# Backtracking state, undo, and result snapshots

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Backtracking explores alternative choices while maintaining a partial state. Define which choices are legal, how one choice changes state, and exactly how to undo it before trying a sibling branch. A visited set may be global for reachability or path-local for a no-reuse path; these are different contracts.

A saved slice header still shares its backing storage. Copy a partial sequence when the result must survive subsequent mutations. Duplicate input values and duplicate output structures are distinct concerns. Output can be exponential or factorial even if each local state update is cheap. The concepts describe exploration mechanics, not a recipe for any listed problem.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
path := []int{4}
path = append(path, 7)
snapshot := append([]int(nil), path...)
path = path[:len(path)-1]
path = append(path, 9)
fmt.Println(path, snapshot)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **[4 9] [4 7]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Write a choose/undo table for selecting one of three colored tokens. Explain when a saved result needs a copy.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Sharing result backing arrays; forgetting undo on early returns; confusing global visited and current-path membership; forgetting output-space cost.

## Questions that use this building block

- [#22 Generate Parentheses](../../22.%20Generate%20Parentheses/README.md)
- [#211 Design Add and Search Words Data Structure](../../211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/README.md)
- [#78 Subsets](../../78.%20Subsets/README.md)
- [#39 Combination Sum](../../39.%20Combination%20Sum/README.md)
- [#46 Permutations](../../46.%20Permutations/README.md)
- [#90 Subsets II](../../90.%20Subsets%20II/README.md)
- [#40 Combination Sum II](../../40.%20Combination%20Sum%20II/README.md)
- [#79 Word Search](../../79.%20Word%20Search/README.md)
- [#131 Palindrome Partitioning](../../131.%20Palindrome%20Partitioning/README.md)
- [#17 Letter Combinations of a Phone Number](../../17.%20Letter%20Combinations%20of%20a%20Phone%20Number/README.md)
- [#51 N-Queens](../../51.%20N-Queens/README.md)

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
