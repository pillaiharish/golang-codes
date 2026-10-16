# LeetCode — Learn Once, Recall Repeatedly

**Purpose:** turn individual Go solutions into **understandable, revisable, repeatable, recallable** mental patterns. Do not learn every problem from scratch before each new interview.

[Start with foundational patterns](0-problem-solving-building-blocks/README.md) · [Rules for new questions](AGENTS.md) · [Track actual revision](REVISION-LOG.md) · [Run all nine problem tests](test-all.sh)

## How to use this index

1. **Identify the pattern** from the question, before opening its implementation.
2. **State the invariant** in a sentence. If it doesn't come, read only that problem's "Invariant" section.
3. **Recreate the algorithm** without copying and dry-run a non-trivial example.
4. **Write Go from memory** and compare with the committed source.
5. **Test** from the individual problem directory: `GO111MODULE=off go test -v` (or `GO111MODULE=off go test -race`).
6. **Repeat all examples** with `bash leetcode/test-all.sh` from the repository root (runs each documented folder with the race detector).
7. **Record uncertainty**, especially which *idea* or Go construct you forgot, in [REVISION-LOG.md](REVISION-LOG.md).
8. **Review deliberately** after roughly 1, 3, 7 and 21 days. If cold recall fails, revise sooner.

### Confidence levels

| Level | What it means |
|---|---|
| 1 | Cannot recognize the pattern |
| 2 | Recognize pattern but cannot outline solution |
| 3 | Explain with notes; struggle to code from memory |
| 4 | Can explain and code correctly without notes |
| 5 | Can adapt to variants, prove invariants, explain trade-offs |

Leave confidence and dates blank until a genuine attempt. **Reading a solution is not the same as recalling it.**

## Pattern quick-reference

| Trigger in problem | Recall this state / invariant | Relevant questions |
|---|---|---|
| Unsorted values + two indices + target | Previously seen value→index; check complement first | [#1](1.%20Two%20Sum/README.md) |
| Any value repeated? | Seen values in a set | [#217](217.%20Contains%20Duplicate/README.md) |
| Same letters in a different order | Compare count signatures | [#242](242.%20Valid%20Anagram/README.md), [#49](49.%20Group%20Anagrams/README.md) |
| Group words with equivalent content | Canonical sorted key or frequency key | [#49](49.%20Group%20Anagrams/README.md) |
| Max in every moving window | Monotonic deque of indices (front = max) | [#239](239.%20Sliding%20Window%20Maximum/README.md) |
| Fast key lookup + least recent eviction | Map + doubly linked list; head MRU, tail LRU | [#146](146.%20LRU%20Cache/README.md) |
| Word/prefix insert and search | Trie: character paths plus isEnd | [#208](208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md) |
| Inorder traversal | Left → Node → Right; recursion or explicit stack | [#94](94.%20Binary%20Tree%20Inorder%20Traversal/README.md) |
| First absent positive in O(n), O(1) space | Index placement: value x belongs at x−1 | [#41](41.%20First%20Missing%20Positive/README.md) |
| Sorted input + target pair | Two pointers from both ends | [Existing #167](167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go) |

## Detailed problem index (LeetCode numbers)

These nine problems use the **same guide structure**: recognition → intuition → invariant → reconstruction recipe → trace → complexity/tradeoffs → cold recall.

| # | Problem guide | Primary idea / related new variant | Solution | Tests | Your confidence | Last recall | Next recall |
|---|---|---|---|---|---|---|---|
| 1 | [Two Sum](1.%20Two%20Sum/README.md) | Complement hash map | [Go](1.%20Two%20Sum/main.go) | [Tests](1.%20Two%20Sum/main_test.go) | — | — | — |
| 41 | [First Missing Positive](41.%20First%20Missing%20Positive/README.md) | In-place cyclic placement | [Go](41.%20First%20Missing%20Positive/main.go) | [Tests](41.%20First%20Missing%20Positive/main_test.go) | — | — | — |
| 49 | [Group Anagrams](49.%20Group%20Anagrams/README.md) | **Existing sorted-key methods preserved**; new array-frequency key | [Original](49.%20Group%20Anagrams/main.go) · [Variant](49.%20Group%20Anagrams/variants.go) | [Tests](49.%20Group%20Anagrams/main_test.go) | — | — | — |
| 94 | [Binary Tree Inorder Traversal](94.%20Binary%20Tree%20Inorder%20Traversal/README.md) | Pointer-to-slice; closure; stack | [Go](94.%20Binary%20Tree%20Inorder%20Traversal/main.go) | [Tests](94.%20Binary%20Tree%20Inorder%20Traversal/main_test.go) | — | — | — |
| 146 | [LRU Cache](146.%20LRU%20Cache/README.md) | Map + doubly linked list | [Go](146.%20LRU%20Cache/main.go) | [Tests](146.%20LRU%20Cache/main_test.go) | — | — | — |
| 208 | [Implement Trie](208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md) | Prefix paths + terminal flags | [Go](208.%20Implement%20Trie%20%28Prefix%20Tree%29/main.go) | [Tests](208.%20Implement%20Trie%20%28Prefix%20Tree%29/main_test.go) | — | — | — |
| 217 | [Contains Duplicate](217.%20Contains%20Duplicate/README.md) | Set (`map[int]struct{}`) | [Go](217.%20Contains%20Duplicate/main.go) | [Tests](217.%20Contains%20Duplicate/main_test.go) | — | — | — |
| 239 | [Sliding Window Maximum](239.%20Sliding%20Window%20Maximum/README.md) | Monotonic deque of indices | [Go](239.%20Sliding%20Window%20Maximum/main.go) | [Tests](239.%20Sliding%20Window%20Maximum/main_test.go) | — | — | — |
| 242 | [Valid Anagram](242.%20Valid%20Anagram/README.md) | **Existing rune maps preserved**; new `[26]int` | [Original](242.%20Valid%20Anagram/main.go) · [Variant](242.%20Valid%20Anagram/variants.go) | [Tests](242.%20Valid%20Anagram/main_test.go) | — | — | — |

## Existing problems still in the repository

These are already present and are **not renamed or rewritten** by this PR. They can be upgraded to the same standard in later PRs. The index should eventually cover *every* existing problem.

| # | Existing work |
|---|---|
| 167 | [Two Sum II — Input Array Is Sorted](167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go) — two pointers, 1-based answers |
| 353 | [Design Snake Game](353%20-%20Design%20Snake%20Game/README.md) |
| 359 | [Logger Rate Limiter](359.%20Logger%20Rate%20Limiter/main.go) |
| 641 | [Design Circular Deque](641.%20Design%20Circular%20Deque/main.go) |
| 757 | [Set Intersection Size At Least Two](757.%20Set%20Intersection%20Size%20At%20Least%20Two/main.go) |
| 1366 | [Rank Teams by Votes](1366.%20Rank%20Teams%20by%20Votes/main.go) |
| 2678 | [Number of Senior Citizens](2678.%20Number%20of%20Senior%20Citizens/main.go) |

## Example: one repeatable study session

Choose **#1 Two Sum**. Before opening code, say:

> I need two indices that add to target. A map stores earlier values with indices. At each number I look for target minus current, **then** insert current.

Run the tests, then close the file and code from memory. Next week, do **#217 Contains Duplicate**. Ask: why is a set enough here, but #1 requires `value→index`? That comparison builds durable recall across problems.

### Questions to ask after every study attempt

- **Pattern gap:** Did I recognize the data structure?
- **Reasoning gap:** Did I state *why* an element can be discarded or moved?
- **Implementation gap:** Did I forget Go map, slice, pointer, or loop syntax?
- **Proof gap:** Could I defend O(n), or did I just repeat it?
- **Transfer gap:** What changes for invalid input, Unicode, sorted data, duplicates, or no extra space?

Log the exact gap; the next revision should target that gap, not blindly reread the entire file.

## Add the next problem

Follow [AGENTS.md](AGENTS.md): keep one numbered folder per problem, preserve previous methods, add working tests and a detailed README, and update this index with a numeric link. The **0 folder is for reusable building blocks**, not a LeetCode problem.
