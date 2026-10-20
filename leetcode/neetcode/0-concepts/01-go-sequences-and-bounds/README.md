# Go sequences, indexing, and ownership

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A slice has a current length and a capacity. Only indices below length are valid. Append returns a new slice header and may replace its backing array. An array has a fixed length as part of its type. Copying a slice shares elements; copying an array copies its elements. A rectangular matrix is normally a slice of separately allocated row slices, not one automatic rectangular allocation.

Before a loop, state the valid index interval. A half-open range [lo,hi) has hi-lo elements. Check emptiness before indexing endpoints. Range yields a copy of each element; changing that copy does not replace an integer element in the slice. Use an index when mutation is intended. Never assume append doubles capacity as a correctness rule.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
s := make([]int, 0, 3)
s = append(s, 4, 8)
view := s
copyOfS := append([]int(nil), s...)
s[0] = 9
fmt.Println(len(s), cap(s), view, copyOfS)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2 3 [9 8] [4 8]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Build a two-row matrix whose rows do not share storage; mutate one row and explain why the other is unchanged.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Indexing spare capacity; discarding append results; unintended shared backing arrays; using the last index as an exclusive bound.

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
- [#125 Valid Palindrome](../../125.%20Valid%20Palindrome/README.md)
- [#167 Two Sum II - Input Array Is Sorted](../../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/README.md)
- [#15 3Sum](../../15.%203Sum/README.md)
- [#11 Container With Most Water](../../11.%20Container%20With%20Most%20Water/README.md)
- [#42 Trapping Rain Water](../../42.%20Trapping%20Rain%20Water/README.md)
- [#121 Best Time to Buy and Sell Stock](../../121.%20Best%20Time%20to%20Buy%20and%20Sell%20Stock/README.md)
- [#3 Longest Substring Without Repeating Characters](../../3.%20Longest%20Substring%20Without%20Repeating%20Characters/README.md)
- [#424 Longest Repeating Character Replacement](../../424.%20Longest%20Repeating%20Character%20Replacement/README.md)
- [#567 Permutation in String](../../567.%20Permutation%20in%20String/README.md)
- [#76 Minimum Window Substring](../../76.%20Minimum%20Window%20Substring/README.md)
- [#239 Sliding Window Maximum](../../239.%20Sliding%20Window%20Maximum/README.md)
- [#20 Valid Parentheses](../../20.%20Valid%20Parentheses/README.md)
- [#155 Min Stack](../../155.%20Min%20Stack/README.md)
- [#150 Evaluate Reverse Polish Notation](../../150.%20Evaluate%20Reverse%20Polish%20Notation/README.md)
- [#22 Generate Parentheses](../../22.%20Generate%20Parentheses/README.md)
- [#739 Daily Temperatures](../../739.%20Daily%20Temperatures/README.md)
- [#853 Car Fleet](../../853.%20Car%20Fleet/README.md)
- [#84 Largest Rectangle in Histogram](../../84.%20Largest%20Rectangle%20in%20Histogram/README.md)
- [#704 Binary Search](../../704.%20Binary%20Search/README.md)
- [#74 Search a 2D Matrix](../../74.%20Search%20a%202D%20Matrix/README.md)
- [#875 Koko Eating Bananas](../../875.%20Koko%20Eating%20Bananas/README.md)
- [#153 Find Minimum in Rotated Sorted Array](../../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/README.md)
- [#33 Search in Rotated Sorted Array](../../33.%20Search%20in%20Rotated%20Sorted%20Array/README.md)
- [#981 Time Based Key-Value Store](../../981.%20Time%20Based%20Key-Value%20Store/README.md)
- [#4 Median of Two Sorted Arrays](../../4.%20Median%20of%20Two%20Sorted%20Arrays/README.md)
- [#287 Find the Duplicate Number](../../287.%20Find%20the%20Duplicate%20Number/README.md)
- [#703 Kth Largest Element in a Stream](../../703.%20Kth%20Largest%20Element%20in%20a%20Stream/README.md)
- [#1046 Last Stone Weight](../../1046.%20Last%20Stone%20Weight/README.md)
- [#973 K Closest Points to Origin](../../973.%20K%20Closest%20Points%20to%20Origin/README.md)
- [#215 Kth Largest Element in an Array](../../215.%20Kth%20Largest%20Element%20in%20an%20Array/README.md)
- [#621 Task Scheduler](../../621.%20Task%20Scheduler/README.md)
- [#355 Design Twitter](../../355.%20Design%20Twitter/README.md)
- [#295 Find Median from Data Stream](../../295.%20Find%20Median%20from%20Data%20Stream/README.md)
- [#78 Subsets](../../78.%20Subsets/README.md)
- [#39 Combination Sum](../../39.%20Combination%20Sum/README.md)
- [#46 Permutations](../../46.%20Permutations/README.md)
- [#90 Subsets II](../../90.%20Subsets%20II/README.md)
- [#40 Combination Sum II](../../40.%20Combination%20Sum%20II/README.md)
- [#79 Word Search](../../79.%20Word%20Search/README.md)
- [#131 Palindrome Partitioning](../../131.%20Palindrome%20Partitioning/README.md)
- [#17 Letter Combinations of a Phone Number](../../17.%20Letter%20Combinations%20of%20a%20Phone%20Number/README.md)
- [#51 N-Queens](../../51.%20N-Queens/README.md)
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
- [#53 Maximum Subarray](../../53.%20Maximum%20Subarray/README.md)
- [#55 Jump Game](../../55.%20Jump%20Game/README.md)
- [#45 Jump Game II](../../45.%20Jump%20Game%20II/README.md)
- [#134 Gas Station](../../134.%20Gas%20Station/README.md)
- [#846 Hand of Straights](../../846.%20Hand%20of%20Straights/README.md)
- [#1899 Merge Triplets to Form Target Triplet](../../1899.%20Merge%20Triplets%20to%20Form%20Target%20Triplet/README.md)
- [#763 Partition Labels](../../763.%20Partition%20Labels/README.md)
- [#678 Valid Parenthesis String](../../678.%20Valid%20Parenthesis%20String/README.md)
- [#57 Insert Interval](../../57.%20Insert%20Interval/README.md)
- [#56 Merge Intervals](../../56.%20Merge%20Intervals/README.md)
- [#435 Non-overlapping Intervals](../../435.%20Non-overlapping%20Intervals/README.md)
- [#252 Meeting Rooms](../../252.%20Meeting%20Rooms/README.md)
- [#253 Meeting Rooms II](../../253.%20Meeting%20Rooms%20II/README.md)
- [#1851 Minimum Interval to Include Each Query](../../1851.%20Minimum%20Interval%20to%20Include%20Each%20Query/README.md)
- [#48 Rotate Image](../../48.%20Rotate%20Image/README.md)
- [#54 Spiral Matrix](../../54.%20Spiral%20Matrix/README.md)
- [#73 Set Matrix Zeroes](../../73.%20Set%20Matrix%20Zeroes/README.md)
- [#202 Happy Number](../../202.%20Happy%20Number/README.md)
- [#66 Plus One](../../66.%20Plus%20One/README.md)
- [#50 Pow(x, n)](../../50.%20Pow%28x%2C%20n%29/README.md)
- [#43 Multiply Strings](../../43.%20Multiply%20Strings/README.md)
- [#2013 Detect Squares](../../2013.%20Detect%20Squares/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/01-arrays-and-slices/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
