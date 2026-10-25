# 981. Time Based Key-Value Store

[NeetCode 150 index](../README.md) · [Concept map](../0-concepts/README.md) · [Official statement](https://leetcode.com/problems/time-based-key-value-store/)

**Category:** Binary Search · **Difficulty:** Medium · **Status:** Unsolved starter

## Problem in your own words

Design timestamped storage supporting set and get for string keys.

Get returns the value at the greatest stored timestamp no later than the query, or an empty string. Set timestamps arrive strictly increasing.

This is an original study summary, not a verbatim copy of LeetCode's statement. Consult the official page for its exact wording, diagrams, and further clarifications. Source metadata was checked on 2026-10-20.

## Input and output contract

```go
func Constructor() TimeMap { /* TODO */ }
func (this *TimeMap) Set(key string, value string, timestamp int)  { /* TODO */ }
func (this *TimeMap) Get(key string, timestamp int) string { /* TODO */ }
```

The function and method signatures above were retrieved from the public LeetCode Go starter. Bodies are intentionally unimplemented. Local node definitions and fixtures replace infrastructure supplied by the judge.

## Constraints and boundary conditions

- 1 <= key.length, value.length <= 100
- key and value consist of lowercase English letters and digits.
- 1 <= timestamp <= 10^7
- All the timestamps timestamp of set are strictly increasing.
- Maximum of 2 * 10^5 calls will be made to set and get .

## Example input and output data

```text
Input: ["TimeMap", "set", "get", "get", "set", "get", "get"] [[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output: [null, null, "bar", "bar", null, "bar2", "bar2"]
```

The example data is for understanding the contract. No explanation, solution trace, or algorithm is supplied. The local demo uses the first source fixture where practical; tree/graph/codec identity cases use small independent fixtures described in `localDemo`.

## Ordered prerequisites

1. [Contracts, complexity, and local test design](../0-concepts/02-contracts-complexity-and-tests/README.md)
2. [Go sequences, indexing, and ownership](../0-concepts/01-go-sequences-and-bounds/README.md)
3. [Binary search bounds and monotonic predicates](../0-concepts/08-binary-search-and-monotonicity/README.md)
4. [Maps, sets, counts, and comparable keys](../0-concepts/03-maps-sets-and-counts/README.md)
5. [Stateful APIs, constructors, and local harnesses](../0-concepts/22-stateful-apis-and-local-harnesses/README.md)

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
