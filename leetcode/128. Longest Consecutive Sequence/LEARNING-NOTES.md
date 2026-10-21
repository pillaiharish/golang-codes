# LeetCode 128 — My attempts, mistakes, and reusable Go/DSA concepts

[← Full problem guide](README.md) · [Correct Go implementation](main.go) · [Regression tests](main_test.go) · [Maps and sets foundation](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md) · [Revision log](../REVISION-LOG.md)

**Purpose:** preserve the actual reasoning path instead of recording only the accepted answer. These are *attempt notes*, not claims of mastery. Revisit the errors before doing the next hash-set or sequence problem.

## The path to the solution

| Stage | Idea | What was right | What needed to change |
|---|---|---|---|
| 1 — Sorting | Sort, deduplicate, compare neighbors for +1, store lengths, take max | Correct sequence definition | O(n log n) sorting misses O(n); a single maximum variable replaces the lengths array |
| 2 — Hash map | Build `map[int]bool` for O(1) membership checks | Sorting isn't necessary for checking numeric neighbors | Naively walking from every input still repeats work |
| 3 — First implementation | Scan **backward** from each original array element; mark `map1[k] = false` to avoid revisits | Recognized that marking visited items could reduce scans | Marking one direction can cut off part of a still-unexplored run; key presence and Boolean truth are different |
| 4 — Second experiment | Remove `map1[k] = false` | Eliminates interference from visited state, which restores the ability to scan complete prefixes | Rescans an entire existing prefix from each new input; worst-case O(n²) and observed time-limit exceeded after 74/84 cases |
| 5 — Correct version | Immutable membership set; only start when `num-1` is **absent**, then count forward | Each distinct number is expanded exactly once across all runs | Expected O(n) time, O(n) memory |

## Attempt 1: sort the values and count neighbors

You first proposed sorting an unsorted input, removing duplicates, checking whether each successive number equals the previous number plus one, saving the lengths of runs and choosing the maximum.

Example: `[100,4,200,1,3,2,2]` becomes `[1,2,3,4,100,200]` after sorting and deduplication; the longest run is `[1,2,3,4]`, length `4`.

**This is a correct algorithm**, provided duplicates do not interrupt a run and a final run is counted after the loop. It costs O(n log n) for sorting, so it cannot meet the requested O(n) complexity. Store only `largest` instead of an entire array of run lengths.

**Reusable concept:** Ask *why* you needed sorting. If the only purpose was to ask whether a value exists, a hash set can provide expected O(1) membership checks without sorting. Go does not need "tuples" to deduplicate numbers: map keys are unique.

## Attempt 2: original failed code (preserved from the study session)

```go
func longestConsecutive(nums []int) int {
	map1 := make(map[int]bool, len(nums))
	for _,v := range nums{
		map1[v] = true
	}
	largest := 0
	for _,j := range nums {
		k := j
		// fmt.Println("k val:",k)
		// fmt.Println("largest val:", largest)
		count := 1
		for map1[k] {
			if _, ok := map1[k-1]; ok{
				// fmt.Println("inside count++")
				map1[k]=false
				count++
				k--
				continue
			}
			if largest < count {
				largest =count
			}
			count = 0
			break
		}
	}
	return largest
}
```

This code is **intentionally shown as a failed historical attempt**, not as a runnable solution or another file in the Go test package. The corrected implementation is in [main.go](main.go).

### Mistake A — "key exists" is not "value is true"

The outer loop asks whether `map1[k]` is `true`; the inner condition `_, ok := map1[k-1]` asks whether a **key exists**, including keys whose value was set to `false`.

```go
m := map[int]bool{2: true}
m[2] = false

fmt.Println(m[2])         // false
_, exists := m[2]
fmt.Println(exists)       // true
fmt.Println(m[3])         // false: key absent, bool zero value
_, absent := m[3]
fmt.Println(absent)       // false

delete(m, 2)              // removes the key entirely
_, exists = m[2]
fmt.Println(exists)       // false
```

**Remember:** Assignment changes a value, not membership. For a map used purely as a set with every stored value `true`, `if seen[x]` conveniently tests membership. If a map also uses `false` to mean "visited", distinguish `present` from `unvisited` explicitly.

### Mistake B — marking as visited destroyed an unfinished run

Trace the actual code on `nums := []int{1,2,3,4}` (expected **4**):

| Outer `j` | What the inner loop does | Map state affected | `largest` |
|---:|---|---|---:|
| `1` | `0` doesn't exist; update maximum with `count=1` | Nothing marked | 1 |
| `2` | Key `1` exists; set `map1[2]=false`, then step to `1`. No `0`; update with `count=2` | `2:false` | 2 |
| `3` | Key `2` **still exists**, so set `map1[3]=false`; next iteration tests `map1[2]`, which is **false**; loop ends without updating maximum | `2:false,3:false` | 2 |
| `4` | Key `3` **still exists**, so set `map1[4]=false`; next iteration tests `map1[3]` false; loop ends without updating maximum | `2:false,3:false,4:false` | 2 |

**Actual result: 2; expected: 4.** This small regression is covered in [main_test.go](main_test.go).

You were decrementing `k`, which explores **only backwards**. Marking `3` visited does not mean the component `[1,2,3,4]` has been fully explored. The algorithm can be dependent on the iteration/input order.

Another failing arrangement is `[3,1,4,2,5]`: starting at `3` consumes `3,2,1`; later processing `4` and `5` cannot connect their runs to the already-consumed middle. The correct answer is `5`, but the first attempt yields `3`.

**General invariant for a visited set:** You may safely mark an item processed only when doing so cannot hide work that later scans still need to perform.

### Mistake C — the maximum is updated only on one exit path

The original code updates `largest` inside the branch for which `k-1` **does not** exist:

```go
if largest < count {
    largest = count
}
count = 0
break
```

But an inner loop may also exit normally when `map1[k]` becomes false. That path never updates `largest`. Resetting `count = 0` just before `break` is redundant because a fresh `count := 1` is created on the next outer iteration.

**Safer loop pattern:** Finish counting one run, then update the maximum **outside the inner loop**. Note that moving the update alone does *not* fix the broken visited-state strategy.

### Mistake D — input values versus distinct map keys

`for _, j := range nums` visits every input element, including duplicate values. `for num := range seen` visits each **unique key** once. Both are valid Go constructs, but the uniqueness guarantee makes the latter easier to reason about for this problem.

Also, `range` over a Go map has unspecified iteration order: do not assume ascending order. The accepted solution works no matter which key is processed first because it uses a mathematical property ("no predecessor") rather than relying on a traversal order.

## Attempt 3: remove visited assignment → time limit exceeded

After removing `map1[k] = false`, an ascending run `[1,2,3,4,5]` is scanned over and over:

| Starting number | Backward sequence inspected | Work |
|---:|---|---:|
| 1 | 1 | 1 |
| 2 | 2, 1 | 2 |
| 3 | 3, 2, 1 | 3 |
| 4 | 4, 3, 2, 1 | 4 |
| 5 | 5, 4, 3, 2, 1 | 5 |

Total for `n` consecutive numbers: `1+2+...+n = n(n+1)/2` → **O(n²)** work. This explains why a valid-looking hash-map solution can still exceed the time limit (74 of 84 cases passed in the experiment).

**Critical complexity lesson:** "Each map lookup is O(1)" is *not* a proof that the entire function is O(n). Count the **total number of lookups** across all loop iterations, not just the cost of a single lookup.

## Correct solution: only start where the predecessor is absent

```go
seen := make(map[int]bool, len(nums))
for _, num := range nums {
    seen[num] = true
}

maxLength := 0
for num := range seen {
    if seen[num-1] {
        continue // This number lies inside an earlier-starting run.
    }

    current := num
    length := 1
    for seen[current+1] {
        current++
        length++
    }
    if length > maxLength {
        maxLength = length
    }
}
return maxLength
```

For `[100,4,200,1,3,2]`:
- `1` has no predecessor (`0`): scan `1 → 2 → 3 → 4` = 4 elements.
- `2,3,4` each have a predecessor: **skip** without scanning.
- `100` and `200` are starts of singleton runs: count 1 each.

The run-start condition guarantees each maximal run is scanned **exactly once**. Though there is a nested loop, the sum of all inner-loop iterations is at most the number of distinct elements, hence **expected O(n)** time with average O(1) map operations and O(n) additional space.

**Why `length := 1`?** A run start is already one number. The inner loop increments only for each *successor* found. You can instead initialize 0 and test the current value first, but be consistent to avoid off-by-one mistakes.

**Why no `false` markers?** The map has one job: membership. Never mutating its values keeps the invariant simple and the code easier to prove.

## Important distinction: visited marking is not inherently wrong

Your instinct to process each element once was useful. A **different** O(n) strategy can start from any unvisited number, mark it visited, and expand **both** backwards and forwards:

```go
// Educational alternative, not the primary solution:
func longestConsecutiveVisited(nums []int) int {
    unvisited := make(map[int]bool, len(nums))
    for _, num := range nums {
        unvisited[num] = true
    }
    largest := 0
    for _, num := range nums {
        if !unvisited[num] {
            continue
        }
        unvisited[num] = false
        count := 1
        for k := num - 1; unvisited[k]; k-- {
            count++
            unvisited[k] = false
        }
        for k := num + 1; unvisited[k]; k++ {
            count++
            unvisited[k] = false
        }
        if count > largest {
            largest = count
        }
    }
    return largest
}
```

Here the entire connected run is consumed before another candidate is processed. Each distinct number becomes false only once. **This alternative works**, but has more mutable state and more failure modes than the sequence-start technique; keep the main [main.go](main.go) version as the interview answer.

There is also a mirror-image O(n) solution: only start at a number with **no successor**, then scan backwards. For that version, iterate **unique map keys**, not duplicate input positions.

## Quick Go foundations learned here

| Syntax | Meaning | Common misunderstanding |
|---|---|---|
| `make(map[int]bool, len(nums))` | Create a writable map, with an initial capacity hint | Not a fixed size, nor guaranteed zero reallocation |
| `seen[num] = true` | Insert key and store Boolean true | Repeated key assignments do not create duplicates |
| `seen[num]` | Read Boolean value; missing key yields false | False does **not** prove a key is absent |
| `_, ok := seen[num]` | Check key existence even if stored value is false | `ok` is separate from stored value |
| `seen[num] = false` | Retain key, change Boolean | Does **not** delete an entry |
| `delete(seen, num)` | Remove key entirely | Unlike marking the key false |
| `for _, num := range nums` | Visit slice elements in slice order | Can process duplicates |
| `for num := range seen` | Visit distinct map keys in unspecified order | Not numerically sorted |
| `continue` | Skip to next iteration of the nearest enclosing loop | Does not exit the whole loop |
| `for seen[current+1]` | Repeat while successor is present-and-true | Nested loop isn't *automatically* quadratic |
| `if length > maxLength` | Keep only best result seen | No need to save all intermediate run lengths |

## Cold recall exercises: re-solve the reasoning, not just the code

1. Without running anything, trace the failed original function on `[1,2,3,4]` and explain precisely **why** it returns `2`.
2. What do `seen[x]` and `_, ok := seen[x]` report for a missing key? For a present key storing `false`?
3. Explain the O(n²) cost of scanning backwards from every number, despite expected O(1) map lookups.
4. Explain the **unique start element** of every maximal consecutive run and why this proves the correct algorithm works even with arbitrary map iteration order.
5. Implement the primary solution cold, then compare it against the two-direction visited alternative. Which invariants are more difficult to maintain?
6. Verify `[]`, `[2,2,2]`, `[3,1,4,2,5]`, negative/zero cases, and one long sorted run.

Record what you actually recalled in [REVISION-LOG.md](../REVISION-LOG.md); do not mark confidence as mastered just because the explanation looks familiar.
