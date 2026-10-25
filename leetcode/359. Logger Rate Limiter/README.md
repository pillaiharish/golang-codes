# 359. Logger Rate Limiter

[← LeetCode study index](../README.md) · [Commented original Go implementation](main.go) · [Runnable tests](main_test.go) · [Quick revision card](REVISION-CARD.md) · [AlgoMonster problem #359](https://algo.monster/liteproblems/359)

**Difficulty:** Easy · **Category:** Design / Hash Table / Data Stream

> **Source note:** This is a faithful, independently worded version of the problem described by AlgoMonster, not a verbatim copy of its copyrighted question or editorial. The six-call example below follows that page's worked example. Additional tests are separately identified as our own.

## Problem statement — read before opening the implementation

Build a logger that decides whether an incoming message is eligible to be printed. Calls supply a timestamp (an integer number of seconds) and the message text.

For **each distinct message**, allow the first occurrence. Once a message is allowed at time `t`, reject the same message at every timestamp **strictly earlier than `t + 10`**. At `t + 10`, it becomes eligible again. Unrelated messages have independent waiting periods.

Implement the following API:

~~~go
type Logger struct {
    // Choose your state representation.
}

func Constructor() Logger

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool
~~~

**Return value:**
- `true` means the caller may print this occurrence.
- `false` means it is still suppressed for that particular message.

**Given behavior / assumptions:**
- Calls arrive in **non-decreasing timestamp order**; you do not need to support time moving backward.
- Different messages may arrive in the same second.
- Several occurrences of the **same** message may arrive in the same second.
- Only successful/allowed messages start a new 10-second waiting period. A rejected call does **not** extend that period.
- `ShouldPrintMessage` decides whether to print; the method does **not** itself write to stdout.

The cited AlgoMonster page provides the timing/order requirements but does **not** publish numerical input-size or message-length bounds in its problem-description section. Do not invent those as official constraints.

## Source example — operations and expected results

The following six calls are the sequence used in the [AlgoMonster walkthrough](https://algo.monster/liteproblems/359#example-walkthrough); the presentation below uses the usual design-problem input/output format.

**Input:**

~~~text
Operations:
["Logger", "shouldPrintMessage", "shouldPrintMessage",
 "shouldPrintMessage", "shouldPrintMessage", "shouldPrintMessage",
 "shouldPrintMessage"]

Arguments:
[[], [1,"foo"], [2,"bar"], [3,"foo"], [8,"bar"],
 [11,"foo"], [12,"bar"]]
~~~

**Expected output:**

~~~text
[null, true, true, false, false, true, true]
~~~

`null` represents the constructor, which produces no boolean result.

| Timestamp | Message | Expected | Reason |
|---:|---|---|---|
| 1 | `foo` | `true` | First `foo`, allowed; next allowed at 11 |
| 2 | `bar` | `true` | First `bar`, allowed; next allowed at 12 |
| 3 | `foo` | `false` | `foo` is waiting until 11 |
| 8 | `bar` | `false` | `bar` is waiting until 12 |
| 11 | `foo` | `true` | Exactly ten seconds since `foo` was allowed |
| 12 | `bar` | `true` | Exactly ten seconds since `bar` was allowed |

**Why the boundary matters:** if `foo` prints at 1, then at 10 it must return `false`; at 11 it must return `true`.

## Extra practice scenarios (not copied from the source)

1. **Repeated arrivals at the same second:** `(5,"ping") → true`, then `(5,"ping") → false`.
2. **A rejected arrival does not reset the timer:** `(1,"ping") → true`, `(9,"ping") → false`, `(11,"ping") → true`.
3. **Separate messages:** `(1,"a") → true`, `(2,"b") → true`, `(10,"a") → false`, `(12,"b") → true`.
4. **New logger state:** a fresh `Logger` instance allows a message even if another instance previously suppressed it.

All these sequences are covered by [`main_test.go`](main_test.go), together with the source example and the exact 9/10/11-second boundary.

### Run tests locally

This repository has independent `package main` problem folders and no root Go module. Run from this problem directory:

~~~bash
cd "leetcode/359. Logger Rate Limiter"
GO111MODULE=off go test -v .
~~~

Or run the existing interactive example:

~~~bash
go run main.go
~~~

## Pattern recognition (after attempting the question)

**Trigger:** a streaming decision based on a *key* (message text) and the time at which that key was previously accepted.

A hash map works because each message needs independent timing state. A simple invariant for the existing implementation is:

> After each call, if a message has been printed, the map stores its **most recent successfully printed timestamp**. Suppressed calls must not update it.

If a message has never been seen, accept it. Otherwise, accept when at least 10 seconds have elapsed since its last accepted timestamp. Only then save the new accepted timestamp.

An equally valid alternative is to store each message's **next permitted timestamp** instead. The existing [Go solution](main.go) uses the last-accepted timestamp, so that code is preserved unchanged.

## Five-step implementation recipe — and why each step exists

1. **Declare** `logs map[string]int` as a field in `Logger` — messages must remember accepted timestamps across multiple method calls. A map declared **inside** `ShouldPrintMessage` would forget everything on the next invocation.
2. **Construct** with `Logger{logs: make(map[string]int)}` — writing to an uninitialized nil map panics. `Constructor() Logger` returns a **value** with a writable map inside it.
3. **Lookup** `previous, exists := this.logs[message]` — distinguishes a never-seen message from an earlier accepted timestamp of zero (a zero map lookup alone cannot distinguish them).
4. **Reject** when `exists && timestamp-previous < 10` — at nine seconds the cooldown holds; at exactly ten it expires. A rejected call must **not** update state.
5. **On acceptance**, store `this.logs[message] = timestamp` and return true — this event, not a suppressed one, begins the next ten-second window.

**Your earlier Go concept gaps:** persistent struct fields versus local variables, constructors and `make`, nil maps versus empty maps, `Logger{}` versus `Constructor()`, and receivers. Reconstruct the same stateful pattern in a simple Counter exercise after you can explain this logger. See the [90-second revision card](REVISION-CARD.md) and [structs/pointers foundation](../0-problem-solving-building-blocks/11-structs-pointers-methods-and-linked-lists/README.md).

**30-second explanation:** I keep one last-accepted timestamp per message in a map owned by the logger instance. On each call I return false if that message was accepted less than ten seconds ago; otherwise I save the current timestamp and return true. Only allowed calls update the state. Hash operations are average O(1) per call, with O(number of unique messages) memory.

## Dry-run the map (existing implementation)

| Call | Stored most-recent print times after call |
|---|---|
| `(1, "foo") → true` | `foo:1` |
| `(2, "bar") → true` | `foo:1, bar:2` |
| `(3, "foo") → false` | unchanged |
| `(8, "bar") → false` | unchanged |
| `(11, "foo") → true` | `foo:11, bar:2` |
| `(12, "bar") → true` | `foo:11, bar:12` |

## Complexity and boundaries

**Average-case time:** `O(1)` per call for a hash-map lookup/update.  
**Additional space:** `O(M)`, where `M` is the number of unique message strings ever encountered.

This interview implementation keeps one entry for each distinct message indefinitely, so storage may grow with new messages. It is **not** a thread-safe or bounded-memory production logger: concurrency protection and stale-key eviction are separate system-design concerns.

## Recall without looking at the solution

1. Why must rejected calls leave the stored timestamp unchanged?
2. What should happen exactly 9, 10 and 11 seconds after an allowed message at time 1?
3. Which map value would you store: last accepted time or next eligible time? Why?
4. Why must `foo` and `bar` maintain separate timestamps?
5. Recreate the six-call source example and implement `Constructor` plus `ShouldPrintMessage` from memory.

**Personal revision status:** Confidence — · Last recalled — · Next recall —

## Original implementation notes

The original `main.go` already implements this specification by storing last-printed timestamps. Its logic is preserved; numbered explanatory comments have been added to support reconstruction from a blank file.
