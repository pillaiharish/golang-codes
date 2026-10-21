# Senior Go concurrency interview preparation

[Senior preparation index](../README.md) · [Repository index](../../README.md)

A dedicated concurrency track for a Go engineer preparing for senior interviews around eight years of experience: **12 topic guides, 60 representative questions, and 14 unsolved implementation tasks**. It develops the research report's emphasis on correctness, production ownership and distributed-systems judgment. Questions are authored practice prompts, not verified lists of questions from particular employers.

## How to use this track

1. Answer a topic's questions aloud before opening the answer outlines.
2. Revisit the building blocks and explain one concrete failure schedule.
3. Implement the linked task, including its cancellation and overload cases.
4. Record actual evidence in [NOTES.md](NOTES.md) and the [senior revision log](../REVISION-LOG.md).
5. Use the [assessment rubric and mock interviews](ASSESSMENT.md) to defend your choices.

The guides explain concepts; the task directories deliberately contain specifications and acceptance criteria rather than completed solutions. Existing numbered modules remain useful reference material, including their worked examples. This track does not renumber or replace them.

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
