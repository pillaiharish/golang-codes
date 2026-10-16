# 0 — Problem-solving building blocks (read before memorizing algorithms)

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
- [Implement Trie #208](../208.%20Implement%20Trie%20(Prefix%20Tree)/README.md).

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
