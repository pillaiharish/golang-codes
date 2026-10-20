# Maps, sets, counts, and comparable keys

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A map associates a comparable key with a value. Use map[T]struct{} for membership, map[T]int for counts, and comma-ok when absence differs from a stored zero. A nil map permits reads but not writes; initialize with make or a literal. Arrays and structs with comparable fields can be keys; slices cannot.

Membership, multiplicity, and original position are different kinds of information. Choose exactly which one your state stores. Map iteration order is not stable, so deterministic presentation requires a separately ordered sequence. Default zero counts make increment convenient but can hide a missing-key distinction. Map operations have expected constant-time behavior in common analysis, not a universal worst-case constant-time guarantee.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
counts := make(map[string]int)
for _, label := range []string{"red", "blue", "red"} { counts[label]++ }
v, ok := counts["green"]
fmt.Println(counts["red"], v, ok)
delete(counts, "blue")
fmt.Println(len(counts))
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2 0 false; then 1**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Track both count and last-seen position for a stream of labels. Explain which keys disappear when you decrement to zero.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Writing to nil maps; assuming iteration order; confusing zero with absence; using []int as a key.

## Questions that use this building block

- [#217 Contains Duplicate](../../217.%20Contains%20Duplicate/README.md)
- [#242 Valid Anagram](../../242.%20Valid%20Anagram/README.md)
- [#1 Two Sum](../../1.%20Two%20Sum/README.md)
- [#49 Group Anagrams](../../49.%20Group%20Anagrams/README.md)
- [#347 Top K Frequent Elements](../../347.%20Top%20K%20Frequent%20Elements/README.md)
- [#238 Product of Array Except Self](../../238.%20Product%20of%20Array%20Except%20Self/README.md)
- [#36 Valid Sudoku](../../36.%20Valid%20Sudoku/README.md)
- [#271 Encode and Decode Strings](../../271.%20Encode%20and%20Decode%20Strings/README.md)
- [#128 Longest Consecutive Sequence](../../128.%20Longest%20Consecutive%20Sequence/README.md)
- [#121 Best Time to Buy and Sell Stock](../../121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/README.md)
- [#3 Longest Substring Without Repeating Characters](../../3.%20Longest%20Substring%20Without%20Repeating%20Characters/README.md)
- [#424 Longest Repeating Character Replacement](../../424.%20Longest%20Repeating%20Character%20Replacement/README.md)
- [#567 Permutation in String](../../567.%20Permutation%20in%20String/README.md)
- [#76 Minimum Window Substring](../../76.%20Minimum%20Window%20Substring/README.md)
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/README.md)
- [#981 Time Based Key-Value Store](../../981.%20Time%20Based%20Key-Value%20Store/README.md)
- [#138 Copy List with Random Pointer](../../138.%20Copy%20List%20with%20Random%20Pointer/README.md)
- [#146 LRU Cache](../../146.%20LRU%20Cache/README.md)
- [#105 Construct Binary Tree from Preorder and Inorder Traversal](../../105.%20Construct%20Binary%20Tree%20from%20Preorder%20and%20Inorder%20Traversal/README.md)
- [#355 Design Twitter](../../355.%20Design%20Twitter/README.md)
- [#200 Number of Islands](../../200.%20Number%20of%20Islands/README.md)
- [#133 Clone Graph](../../133.%20Clone%20Graph/README.md)
- [#695 Max Area of Island](../../695.%20Max%20Area%20of%20Island/README.md)
- [#417 Pacific Atlantic Water Flow](../../417.%20Pacific%20Atlantic%20Water%20Flow/README.md)
- [#130 Surrounded Regions](../../130.%20Surrounded%20Regions/README.md)
- [#994 Rotting Oranges](../../994.%20Rotting%20Oranges/README.md)
- [#286 Walls and Gates](../../286.%20Walls%20and%20Gates/README.md)
- [#207 Course Schedule](../../207.%20Course%20Schedule/README.md)
- [#210 Course Schedule II](../../210.%20Course%20Schedule%20II/README.md)
- [#684 Redundant Connection](../../684.%20Redundant%20Connection/README.md)
- [#323 Number of Connected Components in an Undirected Graph](../../323.%20Number%20of%20Connected%20Components%20in%20an%20Undirected%20Graph/README.md)
- [#261 Graph Valid Tree](../../261.%20Graph%20Valid%20Tree/README.md)
- [#127 Word Ladder](../../127.%20Word%20Ladder/README.md)
- [#332 Reconstruct Itinerary](../../332.%20Reconstruct%20Itinerary/README.md)
- [#1584 Min Cost to Connect All Points](../../1584.%20Min%20Cost%20to%20Connect%20All%20Points/README.md)
- [#743 Network Delay Time](../../743.%20Network%20Delay%20Time/README.md)
- [#778 Swim in Rising Water](../../778.%20Swim%20in%20Rising%20Water/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)
- [#787 Cheapest Flights Within K Stops](../../787.%20Cheapest%20Flights%20Within%20K%20Stops/README.md)
- [#139 Word Break](../../139.%20Word%20Break/README.md)
- [#846 Hand of Straights](../../846.%20Hand%20of%20Straights/README.md)
- [#763 Partition Labels](../../763.%20Partition%20Labels/README.md)
- [#2013 Detect Squares](../../2013.%20Detect%20Squares/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
