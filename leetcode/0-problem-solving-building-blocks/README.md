# 0 — The building-blocks curriculum: Go basics → algorithms → LeetCode

**Additional unsolved practice:** [NeetCode 150 concept-to-problem map](../neetcode/0-concepts/README.md) provides 24 foundation lessons and a prerequisite mapping for all 150 starters.

[← Main LeetCode index](../README.md) · [Study rules](../AGENTS.md) · [Review log](../REVISION-LOG.md) · [Run all foundation examples](verify-examples.py)

Every problem in this repository can be broken down into smaller Go operations and one or two reusable invariants. **Choose either direction:** start with a fundamental, or select a LeetCode problem and study its prerequisites in order. These are **20 specific lessons**, based on the existing problem folders and syntax gaps found during practice, not a generic topic wishlist.

**Debugging-based revision:** [LeetCode 271 Codec attempt — mistake-to-concept map](271-codec-attempt-concept-map.md). This links a real first attempt to the existing Go foundations and new lesson 20; it does not add a purported #271 solution folder.

## How to learn, revise, and recall (the same cycle every time)

1. **Basics first:** open the first prerequisite below; type its complete runnable example, don't just read it.
2. **Explain syntax:** say what each map/slice/loop/pointer operation does, especially boundaries and zero values.
3. **Dry run:** trace the example by hand and predict the result before executing.
4. **Do the mini-drill:** close the lesson and recreate the code. Record which operation or invariant you forgot.
5. **Combine:** go to the LeetCode problem and construct its brute force idea before using the optimized data structure.
6. **Recall on a schedule:** after 1, 3, 7, and ~21 days repeat cold recall, using [REVISION-LOG.md](../REVISION-LOG.md) for real attempts.

**Runnable snippet verification:** from the repository root run `python3 leetcode/0-problem-solving-building-blocks/verify-examples.py`. It extracts the Go fenced program from each lesson and runs it in a temporary directory. A successful run checks compilation/execution of the snippets, **not** correctness of the linked LeetCode solutions. [The nine problem suites](../test-all.sh) are separate.

## A. Foundations by learning order

| # | Foundation lesson (click for code, worked trace and drills) | Operations to practice | LeetCode problems using it |
|---|---|---|---|
| 01 | [Arrays and slices](01-arrays-and-slices/README.md) | Index, swap, append, capacity | [#1](../1.%20Two%20Sum/main.go), [#41](../41.%20First%20Missing%20Positive/main.go), [#94](../94.%20Binary%20Tree%20Inorder%20Traversal/main.go), [#102](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go), [#153](../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go), [#167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), [#200](../200.%20Number%20of%20Islands/main.go), [#239](../239.%20Sliding%20Window%20Maximum/main.go), [#641](../641.%20Design%20Circular%20Deque/main.go), [#1366](../1366.%20Rank%20Teams%20by%20Votes/main.go), [#1460](../1460.%20Make%20Two%20Arrays%20Equal%20by%20Reversing%20Subarrays/main.go), [#1508](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go), [#128](../128.%20Longest%20Consecutive%20Sequence/main.go) |
| 02 | [Loops and bounds](02-loops-and-bounds/README.md) | Range, indices, nested traversals | [#1](../1.%20Two%20Sum/main.go), [#41](../41.%20First%20Missing%20Positive/main.go), [#102](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go), [#153](../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go), [#167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), [#200](../200.%20Number%20of%20Islands/main.go), [#217](../217.%20Contains%20Duplicate/main.go), [#1508](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go), [#2273](../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go), [#2678](../2678.%20Number%20of%20Senior%20Citizens/main.go), [#128](../128.%20Longest%20Consecutive%20Sequence/main.go) |
| 03 | [Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) | Membership, index, counts, buckets | [#1](../1.%20Two%20Sum/main.go), [#49](../49.%20Group%20Anagrams/main.go), [#146](../146.%20LRU%20Cache/main.go), [#217](../217.%20Contains%20Duplicate/main.go), [#242](../242.%20Valid%20Anagram/main.go), [#353](../353%20-%20Design%20Snake%20Game/main.go), [#359](../359.%20Logger%20Rate%20Limiter/main.go), [#1366](../1366.%20Rank%20Teams%20by%20Votes/main.go), [#1460](../1460.%20Make%20Two%20Arrays%20Equal%20by%20Reversing%20Subarrays/main.go), [#128](../128.%20Longest%20Consecutive%20Sequence/main.go) |
| 04 | [Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) | Bytes/runes, `unicode.IsDigit`, ASCII digit scanning, slicing, Atoi | [#49](../49.%20Group%20Anagrams/main.go), [#208](../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go), [#242](../242.%20Valid%20Anagram/main.go), [#1366](../1366.%20Rank%20Teams%20by%20Votes/main.go), [#2273](../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go), [#2678](../2678.%20Number%20of%20Senior%20Citizens/main.go) |
| 05 | [Sorting & comparators](05-sorting-and-comparators/README.md) | Custom priorities and tie-breakers | [#49](../49.%20Group%20Anagrams/main.go), [#757](../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go), [#1366](../1366.%20Rank%20Teams%20by%20Votes/main.go), [#1508](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go), [#2273](../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go) |
| 06 | [Two pointers](06-two-pointers/README.md) | Sorted pairs, invariant, O(1) space | [#167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go) |
| 07 | [Binary search](07-binary-search-and-invariants/README.md) | Monotonic predicate and search bounds | [#153](../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go), [#441](../441.%20Arranging%20Coins/main.go) |
| 08 | [Sliding window + deque](08-sliding-window-and-monotonic-deque/README.md) | Expire, dominate, maintain maximum | [#239](../239.%20Sliding%20Window%20Maximum/main.go) |
| 09 | [Queues, BFS and grids](09-queues-bfs-and-grids/README.md) | FIFO, level size, visited neighbors | [#102](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go), [#200](../200.%20Number%20of%20Islands/main.go), [#353](../353%20-%20Design%20Snake%20Game/main.go) |
| 10 | [Stacks, recursion and DFS](10-stacks-recursion-and-tree-traversal/README.md) | Inorder, recursion and iterative stack | [#94](../94.%20Binary%20Tree%20Inorder%20Traversal/main.go), [#200](../200.%20Number%20of%20Islands/main.go) |
| 11 | [Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) | Pointer receivers, sentinel links | [#94](../94.%20Binary%20Tree%20Inorder%20Traversal/main.go), [#102](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go), [#146](../146.%20LRU%20Cache/main.go), [#208](../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go), [#353](../353%20-%20Design%20Snake%20Game/main.go) |
| 12 | [Circular buffers / deque](12-circular-buffers-and-deques/README.md) | Modulo indices, fixed capacity | [#239](../239.%20Sliding%20Window%20Maximum/main.go), [#641](../641.%20Design%20Circular%20Deque/main.go) |
| 13 | [Tries and 26-letter indexing](13-tries-and-fixed-alphabet-indexing/README.md) | Paths, end markers, array children | [#49](../49.%20Group%20Anagrams/main.go), [#208](../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go), [#242](../242.%20Valid%20Anagram/main.go) |
| 14 | [Greedy interval choice](14-greedy-intervals-and-proof/README.md) | Endpoint sorting and exchange proof | [#757](../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go) |
| 15 | [Subarrays and running sums](15-subarrays-running-sums-and-modular-arithmetic/README.md) | All ranges, prefix sums, modulo | [#441](../441.%20Arranging%20Coins/main.go), [#1508](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go) |
| 16 | [Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) | Transitions, last accepted, moving body | [#146](../146.%20LRU%20Cache/main.go), [#353](../353%20-%20Design%20Snake%20Game/main.go), [#359](../359.%20Logger%20Rate%20Limiter/main.go), [#641](../641.%20Design%20Circular%20Deque/main.go), [#2273](../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go) |
| 17 | [Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Big-O, amortization, preconditions | [#1](../1.%20Two%20Sum/main.go), [#41](../41.%20First%20Missing%20Positive/main.go), [#94](../94.%20Binary%20Tree%20Inorder%20Traversal/main.go), [#146](../146.%20LRU%20Cache/main.go), [#153](../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go), [#167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), [#217](../217.%20Contains%20Duplicate/main.go), [#239](../239.%20Sliding%20Window%20Maximum/main.go), [#441](../441.%20Arranging%20Coins/main.go), [#757](../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go), [#1460](../1460.%20Make%20Two%20Arrays%20Equal%20by%20Reversing%20Subarrays/main.go), [#1508](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go), [#128](../128.%20Longest%20Consecutive%20Sequence/main.go) |
| 18 | [Interfaces and stdlib list](18-interfaces-type-assertions-and-standard-library-containers/README.md) | any, type assertions, container/list | [#353](../353%20-%20Design%20Snake%20Game/main.go) |
| 19 | [In-place cyclic placement](19-in-place-cyclic-placement/README.md) | x belongs at x-1; duplicate guard | [#41](../41.%20First%20Missing%20Positive/main.go) |
| 20 | [Function signatures, types and scope](20-function-signatures-types-and-scope/README.md) | Return types, `:=`/`=`, shadowing, types vs values | [#359](../359.%20Logger%20Rate%20Limiter/main.go), [#271 attempt](271-codec-attempt-concept-map.md) |

## B. Problem → exact prerequisites (all 24 problem folders)

Follow these from **left to right**, and return to the problem once you can state each necessary operation and invariant from memory. Links point to existing Go solutions (including legacy folders whose detailed README is not yet upgraded).

| Existing problem | Read these foundational blocks first | Why these blocks are needed |
|---|---|---|
| [#1 Two Sum](../1.%20Two%20Sum/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Unsorted pair: brute-force loops → seen-value index map |
| [#41 First Missing Positive](../41.%20First%20Missing%20Positive/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [19. In-place cyclic placement](19-in-place-cyclic-placement/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | In-place swaps; first incorrect value-to-index slot |
| [#49 Group Anagrams](../49.%20Group%20Anagrams/main.go) | [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) → [05. Sorting & comparators](05-sorting-and-comparators/README.md) → [13. Tries and 26-letter indexing](13-tries-and-fixed-alphabet-indexing/README.md) | Sorted string signature vs comparable [26]int frequency key |
| [#94 Binary Tree Inorder Traversal](../94.%20Binary%20Tree%20Inorder%20Traversal/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [10. Stacks, recursion and DFS](10-stacks-recursion-and-tree-traversal/README.md) → [11. Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Tree recursion or explicit stack; share the output slice |
| [#102 Binary Tree Level Order Traversal](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [09. Queues, BFS and grids](09-queues-bfs-and-grids/README.md) → [11. Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) | Queue and fixed level size; tree pointer nodes |
| [#128 Longest Consecutive Sequence](../128.%20Longest%20Consecutive%20Sequence/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Membership set; predecessor check to avoid redundant scanning |
| [#146 LRU Cache](../146.%20LRU%20Cache/main.go) | [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [11. Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) → [16. Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Map lookup plus bidirectional recency list |
| [#153 Find Minimum in Rotated Sorted Array](../153.%20Find%20Minimum%20in%20Rotated%20Sorted%20Array/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [07. Binary search](07-binary-search-and-invariants/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Find a binary-search boundary in rotated order |
| [#167 Two Sum II - Input Array Is Sorted](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [06. Two pointers](06-two-pointers/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Sorted two pointers; 1-based indices |
| [#200 Number of Islands](../200.%20Number%20of%20Islands/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [09. Queues, BFS and grids](09-queues-bfs-and-grids/README.md) → [10. Stacks, recursion and DFS](10-stacks-recursion-and-tree-traversal/README.md) | Grid BFS/visited plus bounds and mutation |
| [#208 Implement Trie (Prefix Tree)](../208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go) | [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) → [11. Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) → [13. Tries and 26-letter indexing](13-tries-and-fixed-alphabet-indexing/README.md) | Child arrays, path traversal, terminal flag |
| [#217 Contains Duplicate](../217.%20Contains%20Duplicate/main.go) | [02. Loops and bounds](02-loops-and-bounds/README.md) → [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Set membership is enough; no frequency required |
| [#239 Sliding Window Maximum](../239.%20Sliding%20Window%20Maximum/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [08. Sliding window + deque](08-sliding-window-and-monotonic-deque/README.md) → [12. Circular buffers / deque](12-circular-buffers-and-deques/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Monotonic deque of indices; expiry/dominance |
| [#242 Valid Anagram](../242.%20Valid%20Anagram/main.go) | [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) → [13. Tries and 26-letter indexing](13-tries-and-fixed-alphabet-indexing/README.md) | Rune maps vs fixed lowercase letter counts |
| [#353 Design Snake Game](../353%20-%20Design%20Snake%20Game/main.go) | [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [09. Queues, BFS and grids](09-queues-bfs-and-grids/README.md) → [11. Structs, pointers, linked lists](11-structs-pointers-methods-and-linked-lists/README.md) → [16. Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) → [18. Interfaces and stdlib list](18-interfaces-type-assertions-and-standard-library-containers/README.md) | Linked body, coordinate set, direction transitions |
| [#359 Logger Rate Limiter](../359.%20Logger%20Rate%20Limiter/main.go) | [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [16. Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) | Map of last accepted timestamp per message |
| [#441 Arranging Coins](../441.%20Arranging%20Coins/main.go) | [07. Binary search](07-binary-search-and-invariants/README.md) → [15. Subarrays and running sums](15-subarrays-running-sums-and-modular-arithmetic/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Binary search on triangular-number feasibility |
| [#641 Design Circular Deque](../641.%20Design%20Circular%20Deque/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [12. Circular buffers / deque](12-circular-buffers-and-deques/README.md) → [16. Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) | Ring indices, empty/full conditions, modulo |
| [#757 Set Intersection Size At Least Two](../757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go) | [05. Sorting & comparators](05-sorting-and-comparators/README.md) → [14. Greedy interval choice](14-greedy-intervals-and-proof/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Sort intervals and greedily add rightmost points |
| [#1366 Rank Teams by Votes](../1366.%20Rank%20Teams%20by%20Votes/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) → [05. Sorting & comparators](05-sorting-and-comparators/README.md) | Per-position count vectors and comparator |
| [#1460 Make Two Arrays Equal by Reversing Subarrays](../1460.%20Make%20Two%20Arrays%20Equal%20by%20Reversing%20Subarrays/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [03. Maps, sets, frequency & grouping](03-maps-sets-frequency-and-grouping/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Equal multisets via count increments/decrements |
| [#1508 Range Sum of Sorted Subarray Sums](../1508.%20Range%20Sum%20of%20Sorted%20Subarray%20Sums/main.go) | [01. Arrays and slices](01-arrays-and-slices/README.md) → [02. Loops and bounds](02-loops-and-bounds/README.md) → [05. Sorting & comparators](05-sorting-and-comparators/README.md) → [15. Subarrays and running sums](15-subarrays-running-sums-and-modular-arithmetic/README.md) → [17. Complexity and constraints](17-complexity-amortization-and-input-contracts/README.md) | Enumerate all sums, sort, rank select, modulo |
| [#2273 Find Resultant Array After Removing Anagrams](../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go) | [02. Loops and bounds](02-loops-and-bounds/README.md) → [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) → [05. Sorting & comparators](05-sorting-and-comparators/README.md) → [16. Stateful simulation](16-stateful-simulation-and-sequential-comparison/README.md) | Sorted signature compared only to most recent kept word |
| [#2678 Number of Senior Citizens](../2678.%20Number%20of%20Senior%20Citizens/main.go) | [02. Loops and bounds](02-loops-and-bounds/README.md) → [04. Strings and parsing](04-strings-bytes-runes-and-parsing/README.md) | Fixed byte positions, slice age field, strconv.Atoi |

## C. Worked progression: arrays → maps → Two Sum

**Question:** find two indices in `[2,7,11,15]` whose values add to 9.

**Step 1 — array + loops (brute force):** start with a pair of nested index loops:

```go
func twoSumBrute(nums []int, target int) []int {
    for i := 0; i < len(nums); i++ {
        for j := i + 1; j < len(nums); j++ {
            if nums[i]+nums[j] == target { return []int{i,j} }
        }
    }
    return nil
}
```

Trace: compare (2,7) → 9 → return `[0,1]`. Time O(n²), auxiliary space O(1). This uses [arrays](01-arrays-and-slices/README.md) and [nested loops](02-loops-and-bounds/README.md) only.

**Step 2 — map insight (optimized):** for each value `v`, lookup `target-v` in the map of **already visited** numbers; if found, return its saved index. Else save `v → index`. This uses [hash maps](03-maps-sets-frequency-and-grouping/README.md) and is expected O(n) time with O(n) extra space. See [Two Sum #1](../1.%20Two%20Sum/README.md).

**Step 3 — transfer:** for [Two Sum II #167](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), sorting is already promised: use [two pointers](06-two-pointers/README.md) and O(1) extra space. State *why* the change of assumption changes your choice.

## D. Other repeatable learning routes

- **Counting signature:** [arrays](01-arrays-and-slices/README.md) → [map frequencies](03-maps-sets-frequency-and-grouping/README.md) → [string/rune handling](04-strings-bytes-runes-and-parsing/README.md) → [Valid Anagram #242](../242.%20Valid%20Anagram/README.md) → [Group Anagrams #49](../49.%20Group%20Anagrams/README.md).
- **Queues:** [slices/indices](01-arrays-and-slices/README.md) → [FIFO/BFS](09-queues-bfs-and-grids/README.md) → [Tree Level Order #102](../102.%20Binary%20Tree%20Level%20Order%20Traversal/main.go) → [Number of Islands #200](../200.%20Number%20of%20Islands/main.go).
- **Stateful O(1) designs:** [maps](03-maps-sets-frequency-and-grouping/README.md) → [pointers and links](11-structs-pointers-methods-and-linked-lists/README.md) → [state transitions](16-stateful-simulation-and-sequential-comparison/README.md) → [LRU Cache #146](../146.%20LRU%20Cache/README.md).
- **Indices and windows:** [slices](01-arrays-and-slices/README.md) → [circular buffers](12-circular-buffers-and-deques/README.md) → [monotonic deque](08-sliding-window-and-monotonic-deque/README.md) → [Sliding Window Maximum #239](../239.%20Sliding%20Window%20Maximum/README.md).
- **Reasoning before code:** [index placement](19-in-place-cyclic-placement/README.md) → [amortized analysis](17-complexity-amortization-and-input-contracts/README.md) → [First Missing Positive #41](../41.%20First%20Missing%20Positive/README.md).
- **When the idea is right but Go does not compile:** [function contracts and scope](20-function-signatures-types-and-scope/README.md) → [slices and append](01-arrays-and-slices/README.md) → [struct ownership and state](11-structs-pointers-methods-and-linked-lists/README.md) → [length-prefixed string parsing](04-strings-bytes-runes-and-parsing/README.md) → [the #271 mistake map](271-codec-attempt-concept-map.md).

---

## E. Compact original primer (retained)

The following material is the original summary from the previous PR. It is retained so prior explanations and links remain available; the lesson folders above expand it into runnable examples, operations, drills and cross-problem mapping.

[← LeetCode learning index](../README.md) · [Study rules](../AGENTS.md) · [Revision log](../REVISION-LOG.md)

This is the reusable toolbox behind many solutions. **Do not memorize nine unrelated implementations. Learn a small set of data structures and invariants that generate them.**

## The universal seven-question loop

1. **Recognize:** What word in the problem hints at a known pattern? (“pair,” “duplicate,” “group,” “maximum in every window,” “recently used,” “prefix,” “inorder,” “smallest missing”)
2. **Choose state:** What information must be remembered for the next step? A map, set, counts, deque, doubly linked list, trie node, stack, or array position.
3. **Say invariant:** One sentence describing what is always true after processing each input item.
4. **Write recipe:** 3–6 mechanical steps that reconstruct the implementation without memorizing line numbers.
5. **Dry run:** Track *contents* of state after every meaningful change.
6. **Prove + cost:** Explain why no valid candidate is lost and where work/memory go.
7. **Cold recall:** Close this file, write the invariant first, then the code.

## 1. A hash map remembers information you already saw

```go
indexByValue := make(map[int]int)
for i, x := range nums {
    if j, ok := indexByValue[target-x]; ok {
        return []int{j, i}
    }
    indexByValue[x] = i
}
```

- `value, ok := m[key]`: `ok` distinguishes missing keys from a stored zero-value.
- **Why lookup before insertion?** Prevents using the same array position twice in [Two Sum #1](../1.%20Two%20Sum/README.md).
- **Invariant:** Before position `i`, the map contains only values from previously visited positions.
- Typical expected runtime is O(1) per map operation, not a worst-case guarantee.

## 2. A set stores existence, not a count

```go
seen := make(map[int]struct{}, len(nums))
if _, exists := seen[x]; exists { /* duplicate */ }
seen[x] = struct{}{}
```

- `map[K]struct{}` communicates **membership**; `map[K]int` often represents a **frequency or an index**.
- Empty struct values carry no payload, but maps still have substantial **key, control, bucket/table and allocation overhead**. Do not calculate total map memory as merely value-size times entry count.
- `len(nums)` is a **capacity hint**, not a fixed limit nor a promise of zero allocations. Runtime internals depend on Go version.
- [Contains Duplicate #217](../217.%20Contains%20Duplicate/README.md).

## 3. Frequency counts are a signature for anagrams

- Anagrams have exactly the same character counts.
- For **lowercase a–z only**, `[26]int` gives fixed-sized counts; increment for one string and decrement for the other, then ensure all counts are zero.
- For arbitrary Unicode code points, use `map[rune]int` and `for _, r := range s`; `len(s)` measures bytes, not runes.
- For grouping anagrams, a **sorted string** is another canonical signature. All anagrams sort to the same key.
- **[26]int can itself be a map key**: arrays are comparable. Slices are not valid map keys.
- [Valid Anagram #242](../242.%20Valid%20Anagram/README.md) · [Group Anagrams #49](../49.%20Group%20Anagrams/README.md).

## 4. Slices, indices, and queue behavior

- `append` can reuse the backing array or allocate another; never assume an append is necessarily in place.
- `a = a[1:]` is an **O(1) slice operation**, not a copying shift, but can retain the backing array and cause capacity growth when repeatedly appending. An algorithmic deque with O(k) elements does **not automatically** mean a Go implementation keeps only O(k) allocated memory.
- A **monotonic deque of indices** represents surviving candidates. Store indices so you can tell whether one leaves the window; keep values decreasing from front to back for maximum.
- Each index enters and leaves at most once, which explains amortized O(n) total operations.
- [Sliding Window Maximum #239](../239.%20Sliding%20Window%20Maximum/README.md).

## 5. Map + doubly linked list = fast membership + fast reorder

```go
type node struct {
    key, value int
    prev, next *node
}
```

- Map `key → *node` locates entries in expected O(1).
- Doubly linked list removes/moves **a known node** in O(1), unlike scanning a slice.
- Sentinel `head` / `tail` eliminate empty-list branches.
- **Invariant:** Head-adjacent = most recently used; tail-adjacent = least recently used; map entries and real list nodes match exactly.
- [LRU Cache #146](../146.%20LRU%20Cache/README.md).

## 6. Tries are character-by-character paths

```go
type node struct {
    children [26]*node
    isEnd bool
}
```

- Traverse one edge per letter; create missing nodes during insert; never during search.
- `isEnd` distinguishes a **complete word** from a mere prefix.
- The array representation works **only for lowercase a–z** unless you validate or adapt it.
- [Implement Trie #208](../208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md).

## 7. Trees: recursion gives the traversal order

- Inorder means **Left → Node → Right**.
- Recursive helper carrying a `*[]int` modifies one shared slice header through its pointer.
- A closure capturing `out` is often simpler; iterative traversal uses a stack of `*TreeNode`.
- Stack/recursion space is O(h), where h = tree height (O(n) worst case); output stores O(n).
- [Binary Tree Inorder Traversal #94](../94.%20Binary%20Tree%20Inorder%20Traversal/README.md).

## 8. Cyclic placement: the index is the answer's address

- For `n` values, any integer `x` in `[1,n]` belongs at position `x-1`.
- While the current number is in range and **is not already at its correct target**, swap it there. The duplicate guard prevents infinite swaps.
- Once positions are normalized, the first `nums[i] != i+1` identifies the smallest missing positive.
- This algorithm **modifies the input**. Time O(n), auxiliary space O(1).
- [First Missing Positive #41](../41.%20First%20Missing%20Positive/README.md).

## 9. Go-specific habits to rehearse

| Construct | Recall question |
|---|---|
| `v, ok := m[k]` | Is the map entry absent or is the stored value zero? |
| `make([]T, 0, n)` | Are we creating length zero with capacity n? |
| `make(map[K]V, n)` | Is n a hint or a hard capacity? |
| `[26]int` as key | Why can it be a map key when `[]int` cannot? |
| `s[i]-'a'` | Under what input contract is this safe? |
| `for _, r := range s` | Why are we iterating runes rather than bytes? |
| `p.next.prev = p` | Which two neighbors must be rewired? |
| `nums[i], nums[j] = nums[j], nums[i]` | Why is the duplicate guard necessary? |

## 10. Blank-screen recovery checklist

When you forget the solution, **do not immediately read it**. Write:

```text
Pattern: ______________
State to keep: _________
Invariant after each step: _________
Update rule: ___________
Termination/output: _____
Failure case: __________
```

Only then open the relevant problem README. If you remember the invariant but not the syntax, that's a different learning gap; record it specifically in [REVISION-LOG.md](../REVISION-LOG.md).
