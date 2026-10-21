# Concurrency implementation tasks

[Concurrency index](../README.md) · [Topics](../topics/README.md) · [Assessment rubric](../ASSESSMENT.md)

These are unsolved specifications. No implementation, benchmark result or completed test is claimed. Start with deterministic fake dependencies; add real infrastructure only where the task explicitly calls for integration evidence.

## Task index

1. [Concurrent state and immutable snapshots](01-shared-state-and-snapshots/README.md)
2. [Bounded producer-consumer queue](02-bounded-queue/README.md)
3. [Bounded worker pool](03-bounded-worker-pool/README.md)
4. [Cancellable fan-out and first-success](04-cancellable-fanout/README.md)
5. [Cache loading with shared work](05-cache-and-singleflight/README.md)
6. [Rate, concurrency, and tenant admission](06-rate-and-concurrency-limits/README.md)
7. [Batched asynchronous publisher](07-batched-async-publisher/README.md)
8. [Per-key ordered dispatcher](08-per-key-dispatcher/README.md)
9. [Parallel consumer completion frontier](09-consumer-offset-frontier/README.md)
10. [Rebalance-safe consumer coordinator](10-rebalance-coordinator/README.md)
11. [Idempotent effects and transactional outbox](11-idempotency-and-outbox/README.md)
12. [Graceful service shutdown](12-graceful-shutdown/README.md)
13. [Race, deadlock, and leak diagnosis lab](13-debugging-lab/README.md)
14. [Capstone: telemetry ingestion to Kafka](14-telemetry-kafka-capstone/README.md)

## Local workflow

For each task, create an independent `solution/` module inside its directory so unrelated learning examples do not interfere. Use a Go version installed on your machine and record `go version`; pin external dependencies and commit `go.mod`/`go.sum` when you add them. Keep test helpers and production code separate.

Suggested commands **after implementing a task**, from its `solution/` directory:

```sh
go mod init example.com/concurrency-lab/task01
go test ./...
go test -race ./...
go vet ./...
go test -run TestShutdown -count=20 ./...
go test -bench . -benchmem ./...
```

Run `go mod init` only once, with the appropriate task name. The test and benchmark commands require you to write those tests first; a no-tests result is not evidence. Use `go run .` for a demo after adding `main`; these documentation-only task folders are not currently runnable programs. On macOS, install the toolchain requirements for the race detector if needed. Do not run a repository-wide test expecting unrelated standalone exercises to form one module.

For Kafka integration tasks, document the broker version, Go client/version, topic partitions/replication, consumer group protocol, acknowledgements, retries, idempotence, transaction and buffer settings. Use a disposable local test topic, never production credentials or data. Fake-sink tests prove your local protocol; real-broker tests assess the integration separately.

## Suggested progression

1. Tasks 1–4: shared state, bounded queues, lifecycle and cancellation.
2. Tasks 5–8: shared work, admission, batching and ordering.
3. Tasks 9–11: offsets, reassignment and durable effects.
4. Tasks 12–14: shutdown, diagnosis and design defense.

Advance when you can explain the failure tests, not merely when the happy path runs.
