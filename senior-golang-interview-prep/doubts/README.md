# Doubts: Go fundamentals and runtime behavior

[Senior Go preparation](../README.md) · [Repository](../../README.md)

Seven questions explored individually, progressing from language rules to concurrent system behavior. Read in the order below; each guide contains a direct answer, worked examples, pitfalls, practice, interview checkpoints, and primary sources. Invalid snippets are explicitly labeled and are not included in runnable programs.

## Study plan and topic index

| Order | Guide | What to be able to explain | Suggested practice |
|---|---|---|---|
| 1 | [Declarations, assignment, and memory](01-declarations-assignment-and-memory.md) | `var`, `:=`, `=`, reuse/shadowing, inference, stack/heap and backing storage | Inspect escape diagnostics |
| 2 | [Zero values, make, and new](02-zero-values-make-and-new.md) | Every type's zero value, nil operations, length/capacity, typed nil | Predict values before running |
| 3 | [Bytes, runes, and UTF-8](03-bytes-runes-and-utf8.md) | Code points versus bytes/graphemes, indexing, invalid encodings, memory | Trace byte offsets in multilingual text |
| 4 | [Increments and for loops](04-increments-and-for-loops.md) | Statements versus expressions, continue/post, overflow, loop capture | Predict all loop output |
| 5 | [Buffers, blocking, and backpressure](05-buffers-blocking-and-backpressure.md) | Capacity versus wait policy, combo matrix, nil/closed behavior, Kafka overload | Defend an admission and loss policy |
| 6 | [Network I/O and blocking syscalls](06-network-io-and-blocking-syscalls.md) | G/M/P, poller parking versus blocked thread, deadlines, serial worker bottleneck | Capture a real TCP wait trace |
| 7 | [Slice length, capacity, growth, and limits](07-slice-length-capacity-growth-and-limits.md) | Indexing versus capacity, append allocation, aliasing, explicit bounds, arrays | Predict growth and enforce a batch limit |

Suggested sessions: 1–2 together for declarations/storage; 3–4 for text and syntax prediction; 5–6 for producer-consumer design. Read guide 7 after guide 2, then connect it to guide 5. Spend 30–45 minutes per guide initially, then revisit the checkpoints without notes. These are planning estimates, not a fixed deadline or a claim of mastery.

## Detailed section index

### 1. Variable declarations, initialization, assignment, and memory

- [Direct answer](01-declarations-assignment-and-memory.md#direct-answer)
- [Syntax and scope](01-declarations-assignment-and-memory.md#syntax-and-scope)
- [How short declarations work](01-declarations-assignment-and-memory.md#how-short-declarations-work)
- [Type inference does not mean dynamic typing](01-declarations-assignment-and-memory.md#type-inference-does-not-mean-dynamic-typing)
- [Memory at each level](01-declarations-assignment-and-memory.md#memory-at-each-level)
- [Practice: inspect instead of guessing](01-declarations-assignment-and-memory.md#practice-inspect-instead-of-guessing)
- [Interview checkpoints](01-declarations-assignment-and-memory.md#interview-checkpoints)

### 2. Zero values, nil, make, and new

- [Direct answer](02-zero-values-make-and-new.md#direct-answer)
- [Zero-value operations](02-zero-values-make-and-new.md#zero-value-operations)
- [Slice: nil is already useful](02-zero-values-make-and-new.md#slice-nil-is-already-useful)
- [Map: nil reads and initialized writes](02-zero-values-make-and-new.md#map-nil-reads-and-initialized-writes)
- [Channel: nil differs from unbuffered](02-zero-values-make-and-new.md#channel-nil-differs-from-unbuffered)
- [new versus make](02-zero-values-make-and-new.md#new-versus-make)
- [Typed nil inside an interface](02-zero-values-make-and-new.md#typed-nil-inside-an-interface)
- [Interview checkpoints](02-zero-values-make-and-new.md#interview-checkpoints)

### 3. Bytes, runes, UTF-8, and visible characters

- [Direct answer](03-bytes-runes-and-utf8.md#direct-answer)
- [Comparison](03-bytes-runes-and-utf8.md#comparison)
- [Walk through A, é, and 界](03-bytes-runes-and-utf8.md#walk-through-a-é-and-界)
- [Code points are not grapheme clusters](03-bytes-runes-and-utf8.md#code-points-are-not-grapheme-clusters)
- [Invalid UTF-8](03-bytes-runes-and-utf8.md#invalid-utf-8)
- [Memory and choosing a representation](03-bytes-runes-and-utf8.md#memory-and-choosing-a-representation)
- [Interview checkpoints](03-bytes-runes-and-utf8.md#interview-checkpoints)

### 4. Increment statements and for-loop traps

- [Direct answer](04-increments-and-for-loops.md#direct-answer)
- [Valid and invalid locations](04-increments-and-for-loops.md#valid-and-invalid-locations)
- [for execution order and continue](04-increments-and-for-loops.md#for-execution-order-and-continue)
- [Numeric and concurrency traps](04-increments-and-for-loops.md#numeric-and-concurrency-traps)
- [Loop variables and Go versions](04-increments-and-for-loops.md#loop-variables-and-go-versions)
- [Interview checkpoints](04-increments-and-for-loops.md#interview-checkpoints)

### 5. Bounded buffers, blocking, and producer-consumer policy

- [Direct answer: three separate questions](05-buffers-blocking-and-backpressure.md#direct-answer-three-separate-questions)
- [Combination matrix](05-buffers-blocking-and-backpressure.md#combination-matrix)
- [Three producer APIs](05-buffers-blocking-and-backpressure.md#three-producer-apis)
- [Why default is not inherently a drop](05-buffers-blocking-and-backpressure.md#why-default-is-not-inherently-a-drop)
- [Full/empty/nil/closed behavior](05-buffers-blocking-and-backpressure.md#fullemptynilclosed-behavior)
- [Capacity is not an end-to-end resource bound](05-buffers-blocking-and-backpressure.md#capacity-is-not-an-end-to-end-resource-bound)
- [Apply it to a single-reader telemetry-to-Kafka pipeline](05-buffers-blocking-and-backpressure.md#apply-it-to-a-single-reader-telemetry-to-kafka-pipeline)
- [Interview checkpoints](05-buffers-blocking-and-backpressure.md#interview-checkpoints)

### 6. Goroutine blocking on network I/O versus blocking system calls

- [Direct answer](06-network-io-and-blocking-syscalls.md#direct-answer)
- [G, M, and P](06-network-io-and-blocking-syscalls.md#g-m-and-p)
- [Typical network read on Unix](06-network-io-and-blocking-syscalls.md#typical-network-read-on-unix)
- [A truly blocking syscall](06-network-io-and-blocking-syscalls.md#a-truly-blocking-syscall)
- [Comparison and exceptions](06-network-io-and-blocking-syscalls.md#comparison-and-exceptions)
- [Cancellation and deadlines](06-network-io-and-blocking-syscalls.md#cancellation-and-deadlines)
- [Reproducible observation lab](06-network-io-and-blocking-syscalls.md#reproducible-observation-lab)
- [Why this does not fix a serial Kafka consumer](06-network-io-and-blocking-syscalls.md#why-this-does-not-fix-a-serial-kafka-consumer)
- [Interview checkpoints](06-network-io-and-blocking-syscalls.md#interview-checkpoints)

### 7. Slice length, capacity, growth, and explicit limits

- [Answers to the three scenarios](07-slice-length-capacity-growth-and-limits.md#answers-to-the-three-scenarios)
- [Length and capacity describe different things](07-slice-length-capacity-growth-and-limits.md#length-and-capacity-describe-different-things)
- [Scenario 1: three ways to write the first element](07-slice-length-capacity-growth-and-limits.md#scenario-1-three-ways-to-write-the-first-element)
- [Scenario 2: what happens at the growth boundary](07-slice-length-capacity-growth-and-limits.md#scenario-2-what-happens-at-the-growth-boundary)
- [Language guarantee versus runtime policy](07-slice-length-capacity-growth-and-limits.md#language-guarantee-versus-runtime-policy)
- [Why extra capacity helps](07-slice-length-capacity-growth-and-limits.md#why-extra-capacity-helps)
- [Aliasing and why append returns a slice](07-slice-length-capacity-growth-and-limits.md#aliasing-and-why-append-returns-a-slice)
- [Scenario 3: growth is possible, not infinite](07-slice-length-capacity-growth-and-limits.md#scenario-3-growth-is-possible-not-infinite)
- [Enforce a logical maximum explicitly](07-slice-length-capacity-growth-and-limits.md#enforce-a-logical-maximum-explicitly)
- [Restricting capacity does not prohibit append](07-slice-length-capacity-growth-and-limits.md#restricting-capacity-does-not-prohibit-append)
- [Should you switch to an array?](07-slice-length-capacity-growth-and-limits.md#should-you-switch-to-an-array)
- [Where the design is useful](07-slice-length-capacity-growth-and-limits.md#where-the-design-is-useful)
- [Practice and interview checkpoints](07-slice-length-capacity-growth-and-limits.md#practice-and-interview-checkpoints)

## Runnable examples

This folder is an independent Go module using Go 1.22 language semantics, with only standard-library dependencies. Run commands **from `senior-golang-interview-prep/doubts`**, not the repository root. This leaves existing independent examples/modules unchanged.

```sh
go run ./examples/allocations
go run ./examples/values
go run ./examples/slices
go run ./examples/text
go run ./examples/loops
go run ./examples/queues
go run ./examples/netpoll /tmp/doubts-netpoll.trace
go tool trace /tmp/doubts-netpoll.trace
go vet ./...
go build ./...
```

[Example index](examples/README.md) documents expected observations. Go examples demonstrate valid behavior; compile-failure/panic illustrations in the guides are explicitly identified; the slices demonstration catches its deliberate indexing panic. Runtime traces and compiler diagnostics vary by platform/toolchain. Temporary binaries/traces should stay outside the repository.

## Revision tracking

Record unanswered questions and evidence in [NOTES.md](NOTES.md); record completed reviews in the existing [revision log](../REVISION-LOG.md). Do not mark a topic mastered just because a demonstration runs.

## Links back to the main modules

- [Scheduler](../1-goroutines-and-go-scheduler/README.md)
- [Channels and backpressure](../2-channels-select-and-backpressure/README.md)
- [Synchronization](../3-synchronization-and-race-conditions/README.md)
- [Cancellation and shutdown](../4-context-cancellation-and-shutdown/README.md)
- [Memory and escape analysis](../5-memory-escape-analysis-and-gc/README.md)
- [Benchmarking and debugging](../6-benchmarking-profiling-and-debugging/README.md)
- [Telemetry publisher](../7-bounded-concurrent-telemetry-publisher/README.md)
