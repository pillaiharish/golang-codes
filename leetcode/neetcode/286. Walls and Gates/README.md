# 286. Walls and Gates

[NeetCode 150 index](../README.md) · [Concept map](../0-concepts/README.md) · [Official statement](https://leetcode.com/problems/walls-and-gates/)

**Category:** Graphs · **Difficulty:** Medium · **Status:** Unsolved starter

## Problem in your own words

You are given a grid of rooms, walls, and gates. Cells hold one of three values:

- `-1`: wall; movement through it is forbidden.
- `0`: gate (destination).
- `2147483647` (`INF`): empty room awaiting a distance.

**Modify the grid in place** so that each empty room contains the **minimum number of steps needed to reach any gate**. Each step moves horizontally or vertically to an adjacent cell; diagonal moves do not count. Walls block paths.

Keep gates and walls unchanged. If there is no path from a room to any gate, its value stays `INF`. The Go function does **not** return a matrix; the updated argument is the result.

**Readable source:** [AlgoMonster problem #286](https://algo.monster/liteproblems/286) (reviewed 2026-10-21). This README is a paraphrased learning statement, not copied editorial text.

## Input and output contract

```go
func wallsAndGates(rooms [][]int) { /* TODO */ }
```

**Premium source limitation:** LeetCode exposes the title/API metadata but not its Go starter or statement publicly. This is a conventional local Go adaptation, not a claim of an exact judge signature. Verify the signature in your authorized editor before submission.

## Constraints and boundary conditions

- Only up/down/left/right movement is allowed; each adjacent move costs one.
- Walls `-1` and gates `0` must retain their values.
- Unreachable `INF` cells must stay unchanged.
- Verified numeric grid dimensions were not exposed by the accessible reference.

## Example input and output data

**Example 1 — update the input grid (local fixture)**

```text
Input:
rooms = [[INF, -1,   0, INF],
         [INF, INF, INF,  -1],
         [INF,  -1, INF,  -1],
         [  0,  -1, INF, INF]]

Output (same rooms grid after mutation):
rooms = [[3, -1, 0,  1],
         [2,  2, 1, -1],
         [1, -1, 2, -1],
         [0, -1, 3,  4]]
```

Here `INF = 2147483647`.

**Example 2 — unreachable room (practice)**

```text
Input:  rooms = [[0,-1,INF]]
Output: rooms = [[0,-1,INF]]
```

**Example 3 — gate only (practice)**

```text
Input:  rooms = [[0]]
Output: rooms = [[0]]
```

The Go starter stays unsolved. Reference examples and separately constructed practice cases are labeled; they teach expected behavior without including an algorithm or solution trace.

## Ordered prerequisites

1. [Contracts, complexity, and local test design](../0-concepts/02-contracts-complexity-and-tests/README.md)
2. [Go sequences, indexing, and ownership](../0-concepts/01-go-sequences-and-bounds/README.md)
3. [Maps, sets, counts, and comparable keys](../0-concepts/03-maps-sets-and-counts/README.md)
4. [Stacks, queues, deques, and element lifetime](../0-concepts/07-stacks-queues-and-deques/README.md)
5. [Graph representations, coordinates, and visited state](../0-concepts/13-graphs-and-grids/README.md)
6. [Matrix coordinates, boundaries, and mutation order](../0-concepts/24-matrix-coordinates-and-mutation/README.md)

Study the building blocks, then choose and justify your own approach. The category is a curriculum grouping, not a complete algorithm recipe.

## Run locally on macOS

With Go 1.22 or newer installed, change into this problem directory:

```sh
go run main.go
```

The initial run prints a **TODO** message. Implement only the judge-code section, then rerun. `localDemo` already constructs input and calls your function or a method sequence. Change the fixture and add expected-output checks as you practice. Pointer outputs use Go's diagnostic formatting; inspect node identities/links or add a traversal checker rather than comparing printed addresses.

No external packages, network services, or Go module are required for this command. Each folder is independent, so identical names such as `Constructor` do not collide. To submit, copy only judge declarations/implementations; remove local `main`, fixture helpers, TODO marker references, and node declarations that LeetCode supplies. Keep any algorithm imports needed by your completed solution.

## Your reasoning worksheet — intentionally blank

1. Recognition trigger and assumptions: _write after reading_.
2. Naive idea and preferred idea: _derive both yourself_.
3. Invariant and correctness argument: _your proof_.
4. Plain-language implementation recipe: _your steps_.
5. Hand trace and changing state: _your table_.
6. Time/space analysis and failure boundaries: _your analysis_.
7. Blank-screen recall: _reimplement after hiding your notes_.

## Recall prompts

- What exactly is returned or mutated, and does order or identity matter?
- Which constraints eliminate an otherwise plausible approach?
- What is the smallest valid input, and which boundary case challenges your invariant?
- What state persists between steps or calls, and how will you test it?

Confidence: — · Last attempt: — · Next review: —

Log actual attempts in [REVISION-LOG.md](../REVISION-LOG.md). Running a starter does not mark the problem solved.
