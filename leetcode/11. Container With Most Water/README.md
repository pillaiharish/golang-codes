# 11 — Container With Most Water (Go)

[← LeetCode index](../README.md) · [Primary Go 1.21+ solution](main.go) · [Go 1.20-and-earlier variant](variants_pre_go121.go) · [Brute-force reference](variants.go) · [Table-driven tests](main_test.go) · [Revision card](REVISION-CARD.md) · [Official problem](https://leetcode.com/problems/container-with-most-water/) · [Revision log](../REVISION-LOG.md)

**Pattern:** Converging two pointers, eliminating dominated pairs · **Difficulty:** Medium · **Target:** O(n) time, O(1) auxiliary space.

## Problem statement (paraphrased)

Given `height`, an array of `n` nonnegative integers, two vertical lines are drawn at coordinates `(i, 0)`–`(i, height[i])`. Pick distinct indices `i < j` forming a container with the x-axis. Find the **maximum water area**.

```text
area(i, j) = min(height[i], height[j]) * (j - i)
```

The shorter line limits the water height; the difference of indices is the **width**, not the number of array entries. The container may not be tilted.

**Official constraints:** `2 <= n <= 100000`, `0 <= height[i] <= 10000`. The solution also returns `0` for nil/single-element input, though these are outside LeetCode constraints. The input is never modified.

| Input | Answer | Reason |
|---|---:|---|
| `[1,8,6,2,5,4,8,3,7]` | `49` | indices 1 and 8: `min(8,7) * 7` |
| `[1,1]` | `1` | one unit width and height |
| `[4,3,2,1,4]` | `16` | outer lines: `4 * 4` |
| `[1,2,1]` | `2` | outer lines: `1 * 2` |

## Prerequisites

1. [Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): array indexing and length.
2. [Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): `left < right`, shrinking intervals, safe indices.
3. [Two pointers](../0-problem-solving-building-blocks/06-two-pointers/README.md): discard one candidate endpoint with a proof rather than testing every pair.
4. [Complexity and constraints](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md): O(n²) pairs versus a single O(n) scan.

## 1. Recognize the pattern

The problem asks for the maximum value over **pairs of indices**. The score depends on both the distance (`right-left`) and the smaller endpoint. Since the width shrinks as pointers approach, ask whether one endpoint can be ruled out at each step.

**Memory hook:** Width gets smaller every move → only a taller limiting wall might compensate → move the *shorter* wall.

## 2. Brute force first: inspect every pair

A direct solution chooses every `i` and every `j > i`, computes `min(height[i],height[j]) * (j-i)`, and keeps the maximum. See [`maxAreaBruteForce`](variants.go).

For `[1,8,6,2,5,4,8,3,7]`, indices `1` and `8` give width `7` and limiting height `7`, hence `49`. Brute force checks all `n(n-1)/2` pairs.

**Complexity:** O(n²) time and O(1) extra space. Correct, but too slow for `n` up to 100,000.

## 3. Optimized two-pointer invariant

Start with the widest pair: `left=0` and `right=n-1`. At every iteration:

1. Compute area using `min(height[left],height[right])` and `right-left`.
2. Update `maxWater`.
3. If the left wall is shorter, increment `left`; otherwise decrement `right` (equal-height ties may move either side).

**Why can we discard the shorter side?** Suppose `height[left] <= height[right]` and let the current width be `w = right-left`. Every pair keeping this `left` but choosing some `j < right` has area

```text
min(height[left], height[j]) * (j-left)
    <= height[left] * (j-left)
    <  height[left] * (right-left)
    =  current area
```

So **no pair using the discarded left endpoint** can beat the current area, which we have already considered. It is safe to move left. The other case is symmetric. When heights are equal, either endpoint can be removed safely.

**Invariant:** After each move, every pair discarded so far is unable to beat the best area already recorded; therefore an optimal answer is either already in `maxWater` or can still be found between the current pointers.

Unlike [#167 Two Sum II](../167.%20Two%20Sum%20II%20-%20Input%20Array%20Is%20Sorted/main.go), **#11 does not require the input to be sorted**. The pointer choice follows from the area bound, not from value ordering.

## 4. Full dry run

`height = [1,8,6,2,5,4,8,3,7]`

| Step | Left | Right | Heights | Width | Area | Best | Move |
|---:|---:|---:|---|---:|---:|---:|---|
| 1 | 0 | 8 | 1, 7 | 8 | 8 | 8 | left → 1 |
| 2 | 1 | 8 | 8, 7 | 7 | **49** | **49** | right → 7 |
| 3 | 1 | 7 | 8, 3 | 6 | 18 | 49 | right → 6 |
| 4 | 1 | 6 | 8, 8 | 5 | 40 | 49 | right → 5 (tie) |
| 5 | 1 | 5 | 8, 4 | 4 | 16 | 49 | right → 4 |
| 6 | 1 | 4 | 8, 5 | 3 | 15 | 49 | right → 3 |
| 7 | 1 | 3 | 8, 2 | 2 | 4 | 49 | right → 2 |
| 8 | 1 | 2 | 8, 6 | 1 | 6 | 49 | right → 1 |

At `left == right` the loop stops. **Answer: 49.**

## 5. Reconstruct the Go code — learner's Go 1.21+ solution

The **[primary `maxArea` implementation](main.go)** is now the learner's concise two-pointer solution. Go **1.21 or later** is required, because `min` and `max` became built-in functions in that release. No helper functions or `math` import are needed.

```go
func maxArea(height []int) int {
    area := 0
    n := len(height)
    left, right := 0, n-1
    for left < right {
        area = max(area, min(height[left], height[right])*(right-left))
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
    return area
}
```

**Why this is correct:** `min(height[left], height[right])` finds the limiting wall, `right-left` is the width, and `max(area, candidate)` keeps the largest capacity seen. We always discard the shorter wall using the proof in [section 3](#3-optimized-two-pointer-invariant). This is the **same algorithm and O(n)/O(1) complexity** as the previous, more explicit solution.

### Preserve the original code for Go versions before 1.21

The earlier `main.go` implementation has been moved into **[ `variants_pre_go121.go` ](variants_pre_go121.go)** as `maxAreaLegacy`, with its original pointer and comparison logic preserved. Instead of built-ins, it uses:

```go
h := height[left]
if height[right] < h {
    h = height[right]
}
area := h * (right - left)
if area > maxWater {
    maxWater = area
}
```

| Step | New `maxArea` (Go ≥1.21) | Preserved `maxAreaLegacy` (Go ≤1.20) |
|---|---|---|
| Pick shorter wall | `min(height[left], height[right])` | Explicit `if height[right] < h` comparison |
| Keep maximum | `area = max(area, candidate)` | Explicit `if candidate > maxWater` comparison |
| Choose pointer | Move the shorter height | Identical |
| Complexity | O(n) time / O(1) space | Identical |

**How to submit with Go 1.20 or earlier:** Copy only the `maxAreaLegacy` function from [`variants_pre_go121.go`](variants_pre_go121.go) into LeetCode and rename it to `maxArea`. The full folder's new `main.go` **will not compile on older Go**, because it calls built-in `min` and `max`; the existing test file also uses `slices` (Go 1.21+). The standalone legacy algorithm does not need these features.

**Go syntax trap:** Do not write `height := min(height[left], height[right])` inside the loop: shadowing the input slice named `height` breaks later indexing. Use `h` or directly call `min` in the area expression, as the current solution does.

## 6. Complexity and common mistakes

| Approach | Time | Auxiliary space | Usage |
|---|---|---|---|
| Brute-force pair enumeration | O(n²) | O(1) | Proof/reference/tests |
| **Two pointers with Go 1.21+ built-ins** | **O(n)** | **O(1)** | Current primary solution |
| Two pointers with explicit comparisons | O(n) | O(1) | Same algorithm, older Go compatibility |

Each iteration moves exactly one pointer inward, so there can be at most `n-1` iterations.

- **Move the shorter side**, not the taller one. Reducing width without changing the limiting side cannot improve the area.
- Water height uses **`min`**, never `max` of the two heights.
- Width is `right-left`, **not** `right-left+1`.
- Stop on `left < right`; a single line cannot form a container.
- Equal heights are safe to handle by advancing either pointer; this solution moves `right`.
- No sorting, no extra array, and no input mutation.
- Use the problem's nonnegative-height contract; the algorithm is not defined here for negative heights.

## 7. Interview explanation and recall

> Brute force checks each pair in O(n²). To optimize, I start with the widest possible pair using left and right pointers. I compute the area with the smaller height. If I keep the shorter wall and reduce width, I cannot improve this pair's area, so I can safely discard that wall and move its pointer inward. Each move eliminates one endpoint, giving O(n) time and O(1) extra space.

**Blank-screen questions:** Why is the width `right-left`? Why does the **shorter** pointer move? What if heights tie? Why doesn't #11 require sorted input? What changes in Go 1.21 that lets you use both `min` and `max`? How do you write the same logic in Go 1.20?

Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Log only an actual review attempt in the [revision log](../REVISION-LOG.md).

## Run

From repository root:

```sh
cd 'leetcode/11. Container With Most Water'
GO111MODULE=off go test -count=1 -race -v
GO111MODULE=off go vet .
GO111MODULE=off go run main.go
```

The demo prints `49`. On LeetCode submit only `maxArea`, not `main` or the `fmt` import. On Go 1.20 or earlier, use [the standalone legacy variant](variants_pre_go121.go) and rename its function before submitting; do **not** compile it alongside the new `main.go`. The [brute-force variant](variants.go) is a correctness reference, and tests compare all three implementations.
