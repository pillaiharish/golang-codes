# 261. Graph Valid Tree

[NeetCode 150 index](../README.md) · [Concept map](../0-concepts/README.md) · [Official statement](https://leetcode.com/problems/graph-valid-tree/)

**Category:** Graphs · **Difficulty:** Medium · **Status:** Unsolved starter

## Problem in your own words

A graph has `n` nodes numbered `0` through `n-1`, plus undirected edges `edges[i] = [a,b]`. Return `true` exactly when **all** nodes together form one valid tree.

A tree is **connected** (every vertex can be reached from every other vertex) and has **no cycles**. Both properties matter: a graph with two separate groups fails even if each group has no cycle; a connected graph containing a cycle also fails.

A tree with `n` nodes must have `n-1` edges. Isolated vertices are part of the graph, not silently omitted.

**Readable source:** [AlgoMonster problem #261](https://algo.monster/liteproblems/261) (reviewed 2026-10-21). This README is a paraphrased learning statement, not copied editorial text.

## Input and output contract

```go
func validTree(n int, edges [][]int) bool { /* TODO */ }
```

**Premium source limitation:** LeetCode exposes the title/API metadata but not its Go starter or statement publicly. This is a conventional local Go adaptation, not a claim of an exact judge signature. Verify the signature in your authorized editor before submission.

## Constraints and boundary conditions

- Every vertex in the range `0..n-1` matters.
- Edges are undirected, so `[a,b]` represents a two-way connection.
- A connected acyclic graph is a tree; disconnected or cyclic graphs are not.
- The public description does not verify numerical input limits.

## Example input and output data

**Example 1 — valid tree (local fixture)**

```text
Input:  n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
Output: true
```

**Example 2 — cycle (practice)**

```text
Input:  n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
Output: false
```

**Example 3 — disconnected (practice)**

```text
Input:  n = 4, edges = [[0,1],[2,3]]
Output: false
```

**Single-node boundary (practice):** `n = 1, edges = []` describes a valid tree, so the output is `true`.

The Go starter stays unsolved. Reference examples and separately constructed practice cases are labeled; they teach expected behavior without including an algorithm or solution trace.

## Ordered prerequisites

1. [Contracts, complexity, and local test design](../0-concepts/02-contracts-complexity-and-tests/README.md)
2. [Go sequences, indexing, and ownership](../0-concepts/01-go-sequences-and-bounds/README.md)
3. [Maps, sets, counts, and comparable keys](../0-concepts/03-maps-sets-and-counts/README.md)
4. [Stacks, queues, deques, and element lifetime](../0-concepts/07-stacks-queues-and-deques/README.md)
5. [Graph representations, coordinates, and visited state](../0-concepts/13-graphs-and-grids/README.md)
6. [Dependency order and disjoint-set concepts](../0-concepts/14-topology-and-disjoint-sets/README.md)

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
