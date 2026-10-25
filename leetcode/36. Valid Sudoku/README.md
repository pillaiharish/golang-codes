# 36 — Valid Sudoku (Go)

[← LeetCode index](../README.md) · [Commented solution](main.go) · [Table-driven tests](main_test.go) · [Quick revision card](REVISION-CARD.md) · [LeetCode question](https://leetcode.com/problems/valid-sudoku/) · [Revision log](../REVISION-LOG.md)

**Pattern:** One traversal + three sets of previously seen digits · **Difficulty:** Medium.

## Problem contract and examples

Given a **9×9 board** of bytes, return whether all *filled* cells obey Sudoku's no-repeated-digit rules. Each row, each column and each 3×3 sub-box must not repeat any digit from `'1'` through `'9'`. A `'.'` is empty; the board **need not be solvable or complete**. Inputs contain only those ten permitted symbols.

```go
func isValidSudoku(board [][]byte) bool
```

| Board (nine row strings) | Expected | Reason |
|---|---|---|
| `53..7.... / 6..195... / .98....6. / 8...6...3 / 4..8.3..1 / 7...2...6 / .6....28. / ...419..5 / ....8..79` | `true` | No filled-cell constraint is violated |
| Same board, but first row starts `83..7....` | `false` | An `8` repeats in the upper-left 3×3 box |
| All `.........` rows | `true` | Empty cells are ignored |

The first two are LeetCode's standard illustrative cases, reformatted into row strings. Tests also isolate row-only, column-only, box-only, and bottom-right-box conflicts.

## Before solving: Go foundations

1. [01 — Arrays and slices](../0-problem-solving-building-blocks/01-arrays-and-slices/README.md): `[][]byte`, fixed arrays, and indexing.
2. [02 — Loops and bounds](../0-problem-solving-building-blocks/02-loops-and-bounds/README.md): nested row/column loops and `continue`.
3. [03 — Maps and sets](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md): membership checks and duplicate detection.
4. [04 — Bytes versus runes](../0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md): ASCII `'1'` through `'9'` and integer offsets.

## 1. Recognition trigger

**Question:** Do repeated values violate several overlapping groups? **Trigger:** Make one independent membership tracker for each kind of group, and reject the first repeated membership.

Your initial idea of using a map for row/column/box keys was correct. A single map with properly **tagged composite keys** would work too; the mistake would be using only the digit as a key, because a `5` in two unrelated rows is allowed. Since digits and group counts are bounded to 9, three fixed arrays are especially direct.

## 2. Plain-language intuition

Each filled cell belongs to **exactly one row, one column, and one 3×3 box**. Before registering its digit, ask whether that digit was already registered in any of those groups. If so, the board is invalid. Otherwise mark all three and continue. Do not try to solve the puzzle or compare empty cells.

Use `rows[r][digit]`, `cols[c][digit]`, `boxes[boxIndex][digit]` of type `[9][9]bool`. All elements default to `false`, so no initialization loop or map allocation is needed.

## 3. Invariant: what is true during the scan?

**Before processing `(r,c)`,** each tracker contains precisely the filled digits already encountered in its corresponding row, column, or box. A filled value is illegal if and only if one tracker already contains it. We check *before* marking to avoid treating a cell as a duplicate of itself. If no duplicate is found after all 81 cells, every filled digit obeys all three rules.

## 4. Five reconstruction steps (and why)

1. **Declare** `var rows, cols, boxes [9][9]bool` — nine groups × nine possible digits per rule, with zero-value `false`.
2. **Loop** `r := 0..8`, `c := 0..8`; skip `'.'` — every relevant cell is visited once, while blank cells impose no constraint.
3. **Convert** `digit := int(board[r][c] - '1')` and `boxIndex := (r/3)*3 + c/3` — map symbols to `0..8` and 3×3 groups to `0..8`.
4. **Reject** if `rows[r][digit] || cols[c][digit] || boxes[boxIndex][digit]` — violation in *any* one group is sufficient.
5. **Mark** all three true and continue; return true at the end — future cells must see this occurrence.

### Why the box formula works

Integer division creates row bands `0,1,2` and column bands `0,1,2`:

| Cell `(r,c)` | `r/3` | `c/3` | `(r/3)*3 + c/3` | Box |
|---|---:|---:|---:|---|
| `(0,0)` | 0 | 0 | 0 | Top-left |
| `(1,8)` | 0 | 2 | 2 | Top-right |
| `(4,7)` | 1 | 2 | 5 | Middle-right |
| `(8,8)` | 2 | 2 | 8 | Bottom-right |

## 5. Hand trace: conflict inside a box

Start with an otherwise empty board: `(0,0)='5'`, `(2,2)='5'`.

| Cell | Digit index | Box index | State before check | Action |
|---|---:|---:|---|---|
| `(0,0)='5'` | 4 | 0 | All false | Set `rows[0][4]`, `cols[0][4]`, `boxes[0][4]` |
| `(2,2)='5'` | 4 | 0 | `boxes[0][4]` true | Return `false` even though row and column differ |

## 6. Complexity, boundaries, and alternatives

- **Time:** O(81) for this fixed 9×9 board, usually called O(1) with respect to external input size. Generalized to an N×N board, the traversal is O(N²).
- **Space:** Three `[9][9]bool` arrays = fixed O(1) extra space; no board mutation.
- **Don't confuse byte and Unicode:** `board[r][c]` is a byte, and the question guarantees ASCII digits. Subtracting `'1'` is deliberate; no `strconv.Atoi` or `unicode.IsDigit` is needed.
- **Don't confuse validation with solving:** a partial valid board returns true even if some eventual completion would be impossible.
- **Alternative:** a single `map` of tagged keys like `("row", r, digit)` / `("col", c, digit)` / `("box", boxIndex, digit)` can express the same rules, but uses allocations and more key bookkeeping. A key consisting of the digit alone is incorrect.
- **Input contract:** malformed sizes or symbols are outside this function's promised domain; indexing `digit` requires `'1'`..`'9'`.

## 7. 30-second answer and cold recall

> I scan all filled cells once. I keep separate seen-digit tables for nine rows, nine columns and nine 3×3 boxes. A digit maps to 0–8 with `cell-'1'`; a box maps to 0–8 with `(r/3)*3+c/3`. If any of those three entries was already set, I return false; otherwise I mark them. There are 81 checks and fixed-size state, so O(1) time/space for a 9×9 board.

**Close the code, then answer:**

1. What would go wrong if you used just `map[byte]bool` for the board?
2. Derive the box index for `(7,5)` without looking it up.
3. Why do we check before marking?
4. Can a duplicate digit appear in different boxes and still be valid?
5. Write the nested loops and all three updates from a blank file.

**Personal review:** Confidence `— / 5` · Last cold recall `—` · Next recall `—`. Log only real attempts in [REVISION-LOG.md](../REVISION-LOG.md).

Run locally: `cd 'leetcode/36. Valid Sudoku' && GO111MODULE=off go test -race -v`.
