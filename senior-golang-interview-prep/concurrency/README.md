# Senior Go concurrency interview preparation

[Senior preparation index](../README.md) · [Repository index](../../README.md)

A dedicated concurrency track for a Go engineer preparing for senior interviews around eight years of experience: **12 topic guides, 60 representative questions, 14 unsolved implementation tasks, and 33 staged coding questions with reference answers**. It develops the research report's emphasis on correctness, production ownership and distributed-systems judgment. Questions are authored practice prompts, not verified lists of questions from particular employers.

## How to use this track

1. Answer a topic's questions aloud before opening the answer outlines.
2. Revisit the building blocks and explain one concrete failure schedule.
3. Implement the linked task, including its cancellation and overload cases.
4. Record actual evidence in [NOTES.md](NOTES.md) and the [senior revision log](../REVISION-LOG.md).
5. Use the [assessment rubric and mock interviews](ASSESSMENT.md) to defend your choices.

The guides explain concepts; the task directories deliberately contain specifications and acceptance criteria rather than completed solutions. Existing numbered modules remain useful reference material, including their worked examples. This track does not renumber or replace them.

## Learn-by-solving progression: Q1–Q33

Start with the [study and revision roadmap](STUDY-PLAN.md), then attempt each question **without looking at the answer**. Each numbered folder opens on a question-only README with tasks and scenario prompts. Its adjacent `solution/` directory contains the worked explanation and a Go module with runnable examples. Q3 compares unordered execution with explicit channel ordering; Q4 isolates a deliberate panic from its recovered version.

### Level 1: Goroutine basics

Launch, wait, scheduling, panic.

| # | Question README | Solution walkthrough and source |
|---:|---|---|
| Q1 | [Start one goroutine and wait](1-start-goroutine-and-wait/README.md) | [Explanation](1-start-goroutine-and-wait/solution/README.md) · [main.go](1-start-goroutine-and-wait/solution/main.go) · [sleep variant](1-start-goroutine-and-wait/solution/sleep/main.go) |
| Q2 | [Launch five goroutines and wait](2-start-five-goroutines-and-wait/README.md) | [Explanation](2-start-five-goroutines-and-wait/solution/README.md) · [main.go](2-start-five-goroutines-and-wait/solution/main.go) |
| Q3 | [Execution order is not guaranteed](3-goroutine-execution-order/README.md) | [Explanation](3-goroutine-execution-order/solution/README.md) · [unordered.go](3-goroutine-execution-order/solution/unordered/main.go) · [ordered.go](3-goroutine-execution-order/solution/ordered/main.go) |
| Q4 | [Goroutine panic and recover](4-goroutine-panic-and-recover/README.md) | [Explanation](4-goroutine-panic-and-recover/solution/README.md) · [panic.go](4-goroutine-panic-and-recover/solution/panic/main.go) · [recover.go](4-goroutine-panic-and-recover/solution/recovered/main.go) |

### Level 2: Channel fundamentals

Blocking, buffering, ownership, closure.

| # | Question README | Worked solution & Go code |
|---:|---|---|
| Q5 | [Unbuffered send/receive](5-unbuffered-channel-send-receive/README.md) | [Explanation](5-unbuffered-channel-send-receive/solution/README.md) · [main.go](5-unbuffered-channel-send-receive/solution/main.go) |
| Q6 | [Buffered channel capacity](6-buffered-channel-capacity/README.md) | [Explanation](6-buffered-channel-capacity/solution/README.md) · [main.go](6-buffered-channel-capacity/solution/main.go) |
| Q7 | [Producer-consumer](7-producer-consumer/README.md) | [Explanation](7-producer-consumer/solution/README.md) · [main.go](7-producer-consumer/solution/main.go) |
| Q8 | [Closed channel behavior](8-closed-channel-behavior/README.md) | [Explanation](8-closed-channel-behavior/solution/README.md) · [main.go](8-closed-channel-behavior/solution/main.go) |
| Q9 | [Directional channels](9-directional-channels/README.md) | [Explanation](9-directional-channels/solution/README.md) · [main.go](9-directional-channels/solution/main.go) |
| Q10 | [Deadlocks and lifecycle](10-channel-deadlocks-and-lifecycle/README.md) | [Explanation](10-channel-deadlocks-and-lifecycle/solution/README.md) · [main.go](10-channel-deadlocks-and-lifecycle/solution/main.go) |

### Level 3: Select, timeouts, and cancellation

Multiplexing, deadlines, stop signals.

| # | Question README | Worked solution & Go code |
|---:|---|---|
| Q11 | [Select across channels](11-select-multiple-channels/README.md) | [Explanation](11-select-multiple-channels/solution/README.md) · [main.go](11-select-multiple-channels/solution/main.go) |
| Q12 | [Timeout with time.After](12-timeout-with-time-after/README.md) | [Explanation](12-timeout-with-time-after/solution/README.md) · [main.go](12-timeout-with-time-after/solution/main.go) |
| Q13 | [Context cancellation](13-context-cancellation/README.md) | [Explanation](13-context-cancellation/solution/README.md) · [main.go](13-context-cancellation/solution/main.go) |
| Q14 | [Done channel cancellation](14-done-channel-cancellation/README.md) | [Explanation](14-done-channel-cancellation/solution/README.md) · [main.go](14-done-channel-cancellation/solution/main.go) |
| Q15 | [Nonblocking select/default](15-select-default-nonblocking/README.md) | [Explanation](15-select-default-nonblocking/solution/README.md) · [main.go](15-select-default-nonblocking/solution/main.go) |

### Level 4: Shared state and synchronization

Races, locks, initialization, limits.

| # | Question README | Worked solution & Go code |
|---:|---|---|
| Q16 | [Fix a data race](16-data-race-counter/README.md) | [Explanation](16-data-race-counter/solution/README.md) · [main.go](16-data-race-counter/solution/main.go) |
| Q17 | [Protect a shared map](17-protect-shared-map/README.md) | [Explanation](17-protect-shared-map/solution/README.md) · [main.go](17-protect-shared-map/solution/main.go) |
| Q18 | [Singleton with sync.Once](18-singleton-sync-once/README.md) | [Explanation](18-singleton-sync-once/solution/README.md) · [main.go](18-singleton-sync-once/solution/main.go) |
| Q19 | [errgroup and sibling cancellation](19-errgroup-cancel-on-error/README.md) | [Explanation](19-errgroup-cancel-on-error/solution/README.md) · [main.go](19-errgroup-cancel-on-error/solution/main.go) |
| Q20 | [Semaphore concurrency limit](20-semaphore-limit-concurrency/README.md) | [Explanation](20-semaphore-limit-concurrency/solution/README.md) · [main.go](20-semaphore-limit-concurrency/solution/main.go) |

### Level 5: Core concurrency patterns

Worker pools, fan-in, pipelines, broadcast.

| # | Question README | Worked solution & Go code |
|---:|---|---|
| Q21 | [Worker pool](21-worker-pool/README.md) | [Explanation](21-worker-pool/solution/README.md) · [main.go](21-worker-pool/solution/main.go) |
| Q22 | [Fan-out/fan-in](22-fan-out-fan-in/README.md) | [Explanation](22-fan-out-fan-in/solution/README.md) · [main.go](22-fan-out-fan-in/solution/main.go) |
| Q23 | [Cancellable pipeline](23-cancellable-pipeline/README.md) | [Explanation](23-cancellable-pipeline/solution/README.md) · [main.go](23-cancellable-pipeline/solution/main.go) |
| Q24 | [or-done relay](24-or-done/README.md) | [Explanation](24-or-done/solution/README.md) · [main.go](24-or-done/solution/main.go) |
| Q25 | [Tee to two outputs](25-tee-broadcast-two-outputs/README.md) | [Explanation](25-tee-broadcast-two-outputs/solution/README.md) · [main.go](25-tee-broadcast-two-outputs/solution/main.go) |
| Q26 | [Bridge streams](26-bridge-channel-of-channels/README.md) | [Explanation](26-bridge-channel-of-channels/solution/README.md) · [main.go](26-bridge-channel-of-channels/solution/main.go) |
| Q27 | [Rate limiter](27-rate-limiter/README.md) | [Explanation](27-rate-limiter/solution/README.md) · [main.go](27-rate-limiter/solution/main.go) |
| Q28 | [Pub/sub](28-pub-sub/README.md) | [Explanation](28-pub-sub/solution/README.md) · [main.go](28-pub-sub/solution/main.go) |

### Level 6: Production correctness

Leaks, shutdown, diagnostics, memory model.

| # | Question README | Worked solution & Go code |
|---:|---|---|
| Q29 | [Goroutine leak detection](29-detect-goroutine-leak/README.md) | [Explanation](29-detect-goroutine-leak/solution/README.md) · [main.go](29-detect-goroutine-leak/solution/main.go) |
| Q30 | [Graceful shutdown](30-graceful-shutdown/README.md) | [Explanation](30-graceful-shutdown/solution/README.md) · [main.go](30-graceful-shutdown/solution/main.go) |
| Q31 | [Concurrent URL checker](31-concurrent-url-checker/README.md) | [Explanation](31-concurrent-url-checker/solution/README.md) · [main.go](31-concurrent-url-checker/solution/main.go) |
| Q32 | [Happens-before](32-happens-before/README.md) | [Explanation](32-happens-before/solution/README.md) · [main.go](32-happens-before/solution/main.go) |
| Q33 | [pprof goroutine debugging](33-goroutine-pprof/README.md) | [Explanation](33-goroutine-pprof/solution/README.md) · [main.go](33-goroutine-pprof/solution/main.go) |

### Revision and capstone resources

- [Study plan, minimum starter set, five interview stages, three capstones and final checklist](STUDY-PLAN.md)
- [30 concurrency theory questions with concise answers](THEORY-REVIEW.md)
- [Attempt/evidence log](NOTES.md) and [senior assessment rubric](ASSESSMENT.md)

Run a standard solution from its `solution/` directory with `go run .` or `go run -race .`. Q3 uses `go run ./unordered` and `go run ./ordered`; Q4 uses `go run ./panic` (intentional nonzero exit) and `go run ./recovered`. Q19 needs `go mod tidy` once for `errgroup`. These are reference programs, not claims of comprehensive tests or independent mastery. The [advanced tasks](tasks/README.md) remain intentionally unsolved.

## Topic index

| Topic | Interview focus | Implementation tasks |
|---|---|---|
| [01. Goroutines, scheduling, and lifetime](topics/01-goroutines-scheduler-and-lifetime.md) | G/M/P, I/O blocking, ownership and leaks | [03](tasks/03-bounded-worker-pool/README.md), [04](tasks/04-cancellable-fanout/README.md), [13](tasks/13-debugging-lab/README.md) |
| [02. Memory ordering and data ownership](topics/02-memory-model-and-ownership.md) | Happens-before, aliasing and consistent snapshots | [01](tasks/01-shared-state-and-snapshots/README.md), [02](tasks/02-bounded-queue/README.md), [07](tasks/07-batched-async-publisher/README.md) |
| [03. Mutexes, atomics, and coordination](topics/03-mutexes-atomics-and-coordination.md) | Locks, atomics, WaitGroup, Cond, Once and Pool | [01](tasks/01-shared-state-and-snapshots/README.md), [03](tasks/03-bounded-worker-pool/README.md), [05](tasks/05-cache-and-singleflight/README.md) |
| [04. Channels, select, and queue contracts](topics/04-channels-select-and-queue-contracts.md) | Bounded/blocking policies, nil/closed channels and select | [02](tasks/02-bounded-queue/README.md), [03](tasks/03-bounded-worker-pool/README.md), [12](tasks/12-graceful-shutdown/README.md) |
| [05. Context, deadlines, and errors](topics/05-context-deadlines-and-errors.md) | Deadline budgets, joins, errgroup and uncertain writes | [03](tasks/03-bounded-worker-pool/README.md), [04](tasks/04-cancellable-fanout/README.md), [10](tasks/10-rebalance-coordinator/README.md), [12](tasks/12-graceful-shutdown/README.md) |
| [06. Worker pools, fan-out, and ordering](topics/06-worker-pools-and-patterns.md) | Fan-in/out, singleflight, per-key order and reorder bounds | [03](tasks/03-bounded-worker-pool/README.md), [04](tasks/04-cancellable-fanout/README.md), [05](tasks/05-cache-and-singleflight/README.md), [08](tasks/08-per-key-dispatcher/README.md) |
| [07. Backpressure, rate limits, and capacity](topics/07-backpressure-rate-limits-and-capacity.md) | Rate versus concurrency, capacity, fairness and loss | [02](tasks/02-bounded-queue/README.md), [06](tasks/06-rate-and-concurrency-limits/README.md), [07](tasks/07-batched-async-publisher/README.md), [14](tasks/14-telemetry-kafka-capstone/README.md) |
| [08. Kafka producer concurrency and batching](topics/08-kafka-producer-concurrency.md) | Admission versus acknowledgement, batching and retries | [07](tasks/07-batched-async-publisher/README.md), [08](tasks/08-per-key-dispatcher/README.md), [12](tasks/12-graceful-shutdown/README.md), [14](tasks/14-telemetry-kafka-capstone/README.md) |
| [09. Kafka consumers, offsets, and rebalances](topics/09-kafka-consumer-offsets-and-rebalances.md) | Safe commit frontier, assignment epochs and fencing | [08](tasks/08-per-key-dispatcher/README.md), [09](tasks/09-consumer-offset-frontier/README.md), [10](tasks/10-rebalance-coordinator/README.md), [11](tasks/11-idempotency-and-outbox/README.md) |
| [10. Delivery semantics, idempotency, and retries](topics/10-delivery-idempotency-and-retry.md) | Crash windows, atomic deduplication, outbox and quarantine | [09](tasks/09-consumer-offset-frontier/README.md), [10](tasks/10-rebalance-coordinator/README.md), [11](tasks/11-idempotency-and-outbox/README.md), [14](tasks/14-telemetry-kafka-capstone/README.md) |
| [11. Testing, profiling, and observability](topics/11-testing-profiling-and-observability.md) | Deterministic tests, race detection, profiles and metrics | [01](tasks/01-shared-state-and-snapshots/README.md), [07](tasks/07-batched-async-publisher/README.md), [13](tasks/13-debugging-lab/README.md), [14](tasks/14-telemetry-kafka-capstone/README.md) |
| [12. Graceful shutdown and senior design review](topics/12-shutdown-and-design-review.md) | Admission barriers, drain budgets and operational review | [10](tasks/10-rebalance-coordinator/README.md), [12](tasks/12-graceful-shutdown/README.md), [14](tasks/14-telemetry-kafka-capstone/README.md) |

## Implementation task index

- [01. Concurrent state and immutable snapshots](tasks/01-shared-state-and-snapshots/README.md)
- [02. Bounded producer-consumer queue](tasks/02-bounded-queue/README.md)
- [03. Bounded worker pool](tasks/03-bounded-worker-pool/README.md)
- [04. Cancellable fan-out and first-success](tasks/04-cancellable-fanout/README.md)
- [05. Cache loading with shared work](tasks/05-cache-and-singleflight/README.md)
- [06. Rate, concurrency, and tenant admission](tasks/06-rate-and-concurrency-limits/README.md)
- [07. Batched asynchronous publisher](tasks/07-batched-async-publisher/README.md)
- [08. Per-key ordered dispatcher](tasks/08-per-key-dispatcher/README.md)
- [09. Parallel consumer completion frontier](tasks/09-consumer-offset-frontier/README.md)
- [10. Rebalance-safe consumer coordinator](tasks/10-rebalance-coordinator/README.md)
- [11. Idempotent effects and transactional outbox](tasks/11-idempotency-and-outbox/README.md)
- [12. Graceful service shutdown](tasks/12-graceful-shutdown/README.md)
- [13. Race, deadlock, and leak diagnosis lab](tasks/13-debugging-lab/README.md)
- [14. Capstone: telemetry ingestion to Kafka](tasks/14-telemetry-kafka-capstone/README.md)

See the [task index and local workflow](tasks/README.md) for creating independent Go modules, race tests and integration experiments.

## Prerequisites and existing material

- [Structs, constructors and persistent state](../8-structs-constructors-and-state/README.md): understand instance state and map/slice aliasing before protecting it concurrently.
- [Goroutines and scheduler](../1-goroutines-and-go-scheduler/README.md), [channels and backpressure](../2-channels-select-and-backpressure/README.md), and [synchronization](../3-synchronization-and-race-conditions/README.md).
- [Context and shutdown](../4-context-cancellation-and-shutdown/README.md), [memory and GC](../5-memory-escape-analysis-and-gc/README.md), and [profiling](../6-benchmarking-profiling-and-debugging/README.md).
- [Existing telemetry publisher reference](../7-bounded-concurrent-telemetry-publisher/README.md): compare only after attempting the new capstone independently.
- [Bounded/blocking queue terminology](../doubts/05-buffers-blocking-and-backpressure.md) and [network I/O versus syscalls](../doubts/06-network-io-and-blocking-syscalls.md).

## Source and version policy

Primary references are linked in each guide: the Go specification and package documentation, Apache Kafka, the selected Go client documentation, Kubernetes and Prometheus. Kafka 4.1 links are explicitly versioned references, not a claim that 4.1 is the newest release. Record your actual broker/client versions and group protocol when implementing Kafka tasks. Go runtime details and library APIs can change; check the installed version before relying on an implementation detail.

The supplied research informs topic selection. This track does not repeat unverified release claims or employer-specific interview frequency. All numbers in hypothetical scenarios are illustrative or explicitly identified as scenario input; no benchmark has been run merely by writing this guide.
