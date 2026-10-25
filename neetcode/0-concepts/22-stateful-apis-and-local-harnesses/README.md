# Stateful APIs, constructors, and local harnesses

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

Design tasks expose a constructor and methods rather than one independent function. The object must preserve state between calls, and the method sequence is part of the input. A pointer receiver permits mutation of the same object; constructing a new object for every operation destroys its history. A codec must round-trip the information promised by its contract.

LeetCode normally supplies package context, node definitions, and a driver. The local main.go supplies those pieces but leaves algorithm bodies as TODO panics. Its wrapper catches only its dedicated TODO marker; unexpected panics remain failures. Compilation and a TODO message are not correctness tests. After implementation, compare return values and mutations against expected outcomes and add your own boundary cases.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
type Counter struct { Value int }
c := &Counter{}
c.Value++
c.Value++
fmt.Println(c.Value)
other := &Counter{}
fmt.Println(other.Value)
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **2; 0**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Write an operation sequence with an expected state after every step. Include two independent instances to check for accidentally shared global state.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Reinitializing between calls; unintended global state; copying a struct containing shared pointer fields; submitting local main or duplicate judge node definitions.

## Questions that use this building block

- [#271 Encode and Decode Strings](../../271.%20Encode%20and%20Decode%20Strings/README.md)
- [#981 Time Based Key-Value Store](../../981.%20Time%20Based%20Key-Value%20Store/README.md)
- [#146 LRU Cache](../../146.%20LRU%20Cache/README.md)
- [#297 Serialize and Deserialize Binary Tree](../../297.%20Serialize%20and%20Deserialize%20Binary%20Tree/README.md)
- [#208 Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md)
- [#211 Design Add and Search Words Data Structure](../../211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/README.md)
- [#703 Kth Largest Element in a Stream](../../703.%20Kth%20Largest%20Element%20in%20a%20Stream/README.md)
- [#355 Design Twitter](../../355.%20Design%20Twitter/README.md)
- [#295 Find Median from Data Stream](../../295.%20Find%20Median%20from%20Data%20Stream/README.md)
- [#2013 Detect Squares](../../2013.%20Detect%20Squares/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/16-stateful-simulation-and-sequential-comparison/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
