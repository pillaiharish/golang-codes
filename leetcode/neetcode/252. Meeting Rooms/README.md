# 252. Meeting Rooms

[NeetCode 150 index](../README.md) · [Concept map](../0-concepts/README.md) · [Official statement](https://leetcode.com/problems/meeting-rooms/)

**Category:** Intervals · **Difficulty:** Easy · **Status:** Unsolved starter

## Problem in your own words

Given meeting intervals `intervals[i] = [start, end]`, decide whether **one person** could attend **every** meeting from start to finish. Return `true` only when no pair of meetings overlaps. Input meetings may appear in any order.

An interval finishing at time `t` does **not** conflict with an interval starting at `t`: for example, `[1,3]` and `[3,5]` are compatible. Return a boolean; do not return a count or a reorganized schedule.

**Readable source:** [AlgoMonster problem #252](https://algo.monster/liteproblems/252) (reviewed 2026-10-21). This README is a paraphrased learning statement, not copied editorial text.

## Input and output contract

```go
func canAttendMeetings(intervals [][]int) bool { /* TODO */ }
```

**Premium source limitation:** LeetCode exposes the title/API metadata but not its Go starter or statement publicly. This is a conventional local Go adaptation, not a claim of an exact judge signature. Verify the signature in your authorized editor before submission.

## Constraints and boundary conditions

- Each interval contains its meeting's start and end time.
- Endpoints that only touch are not an overlap; simultaneous active meetings are.
- Do not assume the input is sorted.
- Verified numerical bounds were not available in the accessible description. The empty-list practice case below does not assert that empty input is accepted by the judge.

## Example input and output data

**Example 1 — overlapping meetings (reference)**

```text
Input:  intervals = [[0,30],[5,10],[15,20]]
Output: false
```

The meeting from 0 to 30 conflicts with meetings inside that interval.

**Example 2 — nonoverlapping, unsorted input (practice)**

```text
Input:  intervals = [[7,10],[2,4]]
Output: true
```

**Example 3 — shared endpoint (practice)**

```text
Input:  intervals = [[1,3],[3,5]]
Output: true
```

If an empty list is permitted, `intervals = []` has no conflicts and yields `true`.

The Go starter stays unsolved. Reference examples and separately constructed practice cases are labeled; they teach expected behavior without including an algorithm or solution trace.

## Ordered prerequisites

1. [Contracts, complexity, and local test design](../0-concepts/02-contracts-complexity-and-tests/README.md)
2. [Go sequences, indexing, and ownership](../0-concepts/01-go-sequences-and-bounds/README.md)
3. [Sorting, comparators, and deterministic ordering](../0-concepts/05-sorting-and-order/README.md)
4. [Greedy reasoning and interval boundaries](../0-concepts/18-greedy-and-interval-contracts/README.md)

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
