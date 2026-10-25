# #36 Valid Sudoku — 90-second cold recall card

[Full reasoning](README.md) · [Commented code](main.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md)

**Trigger:** A value must be unique in three overlapping groups: row, column and box.

**Invariant:** Seen tables contain exactly the already visited filled digits for each group.

| Step | Reconstruct | Why |
|---|---|---|
| 1 | `var rows, cols, boxes [9][9]bool` | Fixed 9 groups × 9 digits; defaults to false |
| 2 | Loop `r,c` from 0 through 8 | Visit each cell once |
| 3 | Skip `'.'`; `digit := int(board[r][c]-'1')` | Ignore blanks and map `'1'..'9'` to 0..8 |
| 4 | `box := (r/3)*3+c/3`; check three tables | Detect repeated digit in *any* constrained group |
| 5 | Mark all three true; return true after scan | Future cells can detect this one |

**Trace:** With `'5'` at `(0,0)` and `(2,2)`, both map to box 0 and digit index 4: second occurrence returns false.

**Your earlier thought:** One map can work with tagged row/column/box keys; never key only by digit globally. Separate `[9][9]bool` tables simplify the fixed-size input.

**Traps:** Check *before* marking; `r/3` and `c/3` use integer division; ASCII byte subtraction is valid under the problem contract; a valid partial board need not be complete.

**Cold drill:** Compute box ids `(0,8)`, `(4,7)`, `(8,8)`; implement all three checks from scratch and test a box-only duplicate.

Confidence: `—` · Last cold recall: `—` · Next recall: `—`
