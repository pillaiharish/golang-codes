# Matrix coordinates, boundaries, and mutation order

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A matrix position is a row/column pair. For a rectangular matrix, row bounds use the number of rows and column bounds use the row width. Transposing exchanges coordinates and swaps dimensions; only a square matrix keeps the same dimensions. Four boundary indices can describe an inner rectangle, and degenerate one-row or one-column rectangles need explicit treatment.

In-place updates can overwrite information that later steps still need. Separate the original-input condition from mutations caused by that condition. A swap preserves both values, while sequential assignment may lose one. Rows are distinct slices in Go and can technically have different lengths unless the contract promises a rectangle. Changing a matrix in place does not mean that every useful temporary variable is forbidden.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
matrix := [][]int{{1, 2, 3}, {4, 5, 6}}
r, c := 1, 2
fmt.Println(len(matrix), len(matrix[0]), matrix[r][c])
matrix[0][0], matrix[1][0] = matrix[1][0], matrix[0][0]
fmt.Println(matrix)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2 3 6; [[4 2 3] [1 5 6]]**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Label the cells of a 2-by-3 rectangle and its transpose. Explain which information an in-place write might destroy.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Using row count for a column bound; revisiting a degenerate boundary; treating newly changed cells as original input; sharing row backing storage.

## Questions that use this building block

- [#36 Valid Sudoku](../../36.%20Valid%20Sudoku/README.md)
- [#74 Search a 2D Matrix](../../74.%20Search%20a%202D%20Matrix/README.md)
- [#212 Word Search II](../../212.%20Word%20Search%20II/README.md)
- [#79 Word Search](../../79.%20Word%20Search/README.md)
- [#200 Number of Islands](../../200.%20Number%20of%20Islands/README.md)
- [#695 Max Area of Island](../../695.%20Max%20Area%20of%20Island/README.md)
- [#417 Pacific Atlantic Water Flow](../../417.%20Pacific%20Atlantic%20Water%20Flow/README.md)
- [#130 Surrounded Regions](../../130.%20Surrounded%20Regions/README.md)
- [#994 Rotting Oranges](../../994.%20Rotting%20Oranges/README.md)
- [#286 Walls and Gates](../../286.%20Walls%20and%20Gates/README.md)
- [#778 Swim in Rising Water](../../778.%20Swim%20in%20Rising%20Water/README.md)
- [#329 Longest Increasing Path in a Matrix](../../329.%20Longest%20Increasing%20Path%20in%20a%20Matrix/README.md)
- [#48 Rotate Image](../../48.%20Rotate%20Image/README.md)
- [#54 Spiral Matrix](../../54.%20Spiral%20Matrix/README.md)
- [#73 Set Matrix Zeroes](../../73.%20Set%20Matrix%20Zeroes/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../0-problem-solving-building-blocks/01-arrays-and-slices/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
