# 238 — Product of Array Except Self (Go)

[← LeetCode index](../README.md) · [Commented solution](main.go) · [Table-driven tests](main_test.go) · [Quick revision card](REVISION-CARD.md) · [LeetCode question](https://leetcode.com/problems/product-of-array-except-self/) · [Revision log](../REVISION-LOG.md)

**Pattern:** Prefix products + suffix products · **Difficulty:** Medium.

## Problem contract and examples

For each index `i`, return the product of **all input elements except `nums[i]`**. Achieve O(n) time **without division** and O(1) auxiliary space when the returned answer is excluded. The standard problem has at least two numbers and guarantees the relevant prefix/suffix products fit a 32-bit integer.

```go
func productExceptSelf(nums []int) []int
```

| Input | Expected result | Point to notice |
|---|---|---|
| `[1,2,3,4]` | `[24,12,8,6]` | Each index excludes itself |
| `[-1,1,0,-3,3]` | `[0,0,9,0,0]` | Zero makes division unsuitable |
| `[0,2,0,4]` | `[0,0,0,0]` | Two zeros kill every product |

The first two are standard LeetCode examples; the third is a regression case.

## Prerequisite building blocks

1. [01 — Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): allocating and writing the answer.
2. [02 — Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): forwards and backwards indexing.
3. [15 — Running aggregates](../0-problem-solving-building-blocks/15-subarrays-running-sums-and-modular-arithmetic/README.md): products as rolling state, not stored subarrays.
4. [17 — Complexity](../0-problem-solving-building-blocks/17-complexity-amortization-and-input-contracts/README.md): why two independent sweeps remain O(n).

## 1. Recognition trigger

**Question:** Every output depends on all elements *except one*, without division. **Trigger:** Split the product around `i` into what lies strictly to the **left** and strictly to the **right**. You do not need to multiply all other indices separately.

## 2. Intuition: from repeated multiplication to two passes

Your earlier thought of multiplying neighboring indices in successive passes notices that products can be reused. The risk is repeatedly processing nearly the full array for each destination: this can become O(n²). Instead use two running aggregates: a prefix product while moving right and a suffix product while moving left. Write the left product into `ans[i]`, then multiply the right product into that same slot.

The **identity `1`** is vital. At index 0 there are no elements on the left, so left product is 1. At the last index there are no elements on the right, so right product is also 1.

## 3. Invariants / correctness

- **Forward pass:** immediately before writing `ans[i]`, `prefix` equals `nums[0] * ... * nums[i-1]`. Store it *before* `prefix *= nums[i]` so `nums[i]` is excluded.
- **Backward pass:** immediately before updating `ans[i]`, `suffix` equals `nums[i+1] * ... * nums[n-1]`. Multiply `ans[i]` by that *before* `suffix *= nums[i]`.

After both passes, `ans[i] = (all elements to the left) × (all elements to the right)` — precisely the elements other than `i`. Zeros require no special branch.

## 4. Five steps to reconstruct (with reasons)

1. **Allocate** `ans := make([]int, len(nums))` — we may reuse the required output as prefix storage.
2. **Set** `prefix := 1` — the product of an empty range is the multiplicative identity.
3. **Sweep forward:** `ans[i] = prefix; prefix *= nums[i]` — store the exclusive left product first.
4. **Set** `suffix := 1` and **sweep backward:** `ans[i] *= suffix; suffix *= nums[i]` — multiply by the exclusive right product first.
5. **Return** `ans` — one output slice, two running integers; no divisions, nested scans, or extra prefix/suffix arrays.

## 5. Hand trace: `[1,2,3,4]`

| Index | `nums[i]` | `ans[i]` after left sweep | Running prefix after update |
|---:|---:|---:|---:|
| 0 | 1 | 1 | 1 |
| 1 | 2 | 1 | 2 |
| 2 | 3 | 2 | 6 |
| 3 | 4 | 6 | 24 |

| Index (reverse) | Right product before update | Multiply existing `ans[i]` | Result |
|---:|---:|---|---:|
| 3 | 1 | `6 × 1` | 6 |
| 2 | 4 | `2 × 4` | 8 |
| 1 | 12 | `1 × 12` | 12 |
| 0 | 24 | `1 × 24` | 24 |

The return value is `[24,12,8,6]`.

## 6. Complexity and common failures

- O(n) time: the two sweeps make `2n` iterations, not O(n²).
- O(1) **auxiliary** space: just `prefix`, `suffix`, loop indices, **excluding the O(n) answer required by the API**. Including output storage, memory is O(n).
- Updating the running product **before** storing it accidentally includes `nums[i]` (off-by-one error).
- Starting `prefix` or `suffix` at `0` makes every answer zero. Multiplicative identity is `1`.
- A zero product shortcut or division by `nums[i]` is not a general solution: division is disallowed and zero cannot be divided away.
- Do not mutate `nums`; the provided code writes only to `ans`.
- A full `prefix[]` plus `suffix[]` solution is easy to derive and also O(n), but takes O(n) auxiliary space. Reusing `ans` removes those arrays.

## 7. Interview explanation and blank-screen recall

> The product excluding index i is its prefix product times its suffix product. I store each exclusive prefix into the answer on a forward sweep, then walk backward carrying an exclusive suffix and multiply it into each answer slot. Updating each running product *after* using it prevents including the current element. Two O(n) passes, no division, O(1) extra space beyond the output.

**Before reopening code:**

1. Why must `ans[i] = prefix` happen before multiplying `nums[i]`?
2. Why do both accumulators start at `1`, including boundaries?
3. Trace `[1,0,3,4]` and `[0,2,0,4]` without any special zero handling.
4. Why is the space O(1) auxiliary, yet the returned output O(n)?
5. Recreate the two loops with the correct index direction.

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Record actual attempts in [REVISION-LOG.md](../REVISION-LOG.md).

Run locally: `cd 'leetcode/238. Product of Array Except Self' && GO111MODULE=off go test -race -v`.
