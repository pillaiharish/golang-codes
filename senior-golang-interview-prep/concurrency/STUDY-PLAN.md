# Concurrency learn-by-solving roadmap

[Concurrency index](README.md) · [Theory Q&A](THEORY-REVIEW.md) · [Senior assessment](ASSESSMENT.md) · [Revision evidence](NOTES.md)

This roadmap turns 33 small-to-production exercises into a **spiral learning track**. Each numbered folder opens with a **question-only README**; its adjacent `solution/` directory holds the reference code and worked explanation. Begin with recall and a clean implementation before opening that solution. This is study material, not a claim about questions asked by a particular employer.

## How to solve every question

1. Read the task and **predict** which statements can run first, block, panic, race or leak.
2. Write the smallest standalone program without looking at the reference.
3. Execute it normally and with the Go race detector where applicable:

   ```sh
   go run -race main.go
   go test -race ./...
   ```

   Run from a question's `solution/` module. For Q3 use `go run ./unordered` or `go run ./ordered`; for Q4 use `go run ./panic` (expected failure) or `go run ./recovered`. Q19 requires `go mod tidy` once to fetch its pinned `errgroup` dependency. The reference modules include runnable examples, but not a claim that comprehensive tests have been written. Run `go test` after adding tests for the relevant package.
4. Break one contract intentionally: omit a signal, stop receiving, close early, cancel, or introduce an unsafe shared write.
5. Fix it and demonstrate the **reason**, not just a different observed output. Not every execution reveals a race or alternate ordering.
6. State the invariant, ownership, blocking points, failure policy and resource bound. Record what you actually ran in [NOTES.md](NOTES.md), then revisit after 1, 3, 7 and 14 days.

## Level progression

| Level | Questions | Outcome |
|---|---|---|
| **1 — Goroutine basics** | [Q1](1-start-goroutine-and-wait/README.md), [Q2](2-start-five-goroutines-and-wait/README.md), [Q3](3-goroutine-execution-order/README.md), [Q4](4-goroutine-panic-and-recover/README.md) | Start, join, reason about scheduling and panic containment |
| **2 — Channel fundamentals** | [Q5–Q10](README.md#level-2-channel-fundamentals) | Rendezvous, buffers, closure, direction, deadlock/lifecycle |
| **3 — Select and cancellation** | [Q11–Q15](README.md#level-3-select-timeouts-and-cancellation) | Multiplexing, deadlines, cancellation and nonblocking policies |
| **4 — Shared state** | [Q16–Q20](README.md#level-4-shared-state-and-synchronization) | Race safety, maps, once, grouped errors, concurrency bounds |
| **5 — Patterns** | [Q21–Q28](README.md#level-5-core-concurrency-patterns) | Worker pools, fan-in, pipeline, bridge, broadcast and pacing |
| **6 — Production correctness** | [Q29–Q33](README.md#level-6-production-correctness) | Leak diagnosis, shutdown, HTTP concurrency, memory model, profiling |

## Minimum starter set when time is short

1. **WaitGroup and loop variables:** Q1, Q2, Q3.
2. **Channel send/receive, buffering and close:** Q5–Q9.
3. **Deadlock and `select`:** Q10–Q12.
4. **Cancellation:** Q13, Q14.
5. **Data races:** Q16, Q17; contrast mutex versus atomic.
6. **Bounded concurrency:** Q20, Q21.
7. **Fan-out/fan-in:** Q22 and pipeline Q23.
8. **Goroutine leak:** Q29.

Do not mark these complete because a reference answer exists. Prioritize the ones you cannot code or explain from memory.

## Interview preparation stages

| Stage | Practice | Interview evidence |
|---|---|---|
| Syntax and basic behavior | Q1–Q10 | Predict blocked and panic cases; no sleep-based synchronization |
| Coordination | Q11–Q15 plus Q21 | Explain cancellation, ownership and who closes each channel |
| Shared state | Q16–Q20 and Q32 | Identify happens-before, protect invariants and bound admission |
| Patterns | Q21–Q28 | State ordering, resource, error and shutdown contracts |
| Production | Q29–Q33 | Reproduce leaks/deadlocks, profile, test and defend trade-offs |

## Three larger capstones

### Capstone A — Concurrent file processor

Input: a list of paths. Read/process files concurrently with a limit N. Collect per-file results and errors; cancel on parent deadline, stop admission, join workers, and shut down cleanly.

**Acceptance:** deterministic max-active-N test; stable input-to-output identity; unreadable-file failure; cancellation while waiting for admission; no blocked sender after caller stops; explicit memory and open-file bounds.

**Relevant questions:** [Q13](13-context-cancellation/README.md), [Q20](20-semaphore-limit-concurrency/README.md), [Q21](21-worker-pool/README.md), [Q30](30-graceful-shutdown/README.md).

### Capstone B — Concurrent task queue

Support multiple producers, bounded admission, multiple workers and structured results. Classify completed, failed, canceled-before-start and unresolved-at-shutdown jobs. A critical failure may cancel siblings, but all accepted work needs accounting.

**Acceptance:** concurrent Submit/Shutdown; queue-full behavior; result draining; first failure cancellation; per-job completion bookkeeping; no send-on-closed panic; test a worker that ignores context.

**Relevant questions:** [Q15](15-select-default-nonblocking/README.md), [Q19](19-errgroup-cancel-on-error/README.md), [Q21](21-worker-pool/README.md), [Q30](30-graceful-shutdown/README.md), [advanced task 03](tasks/03-bounded-worker-pool/README.md).

### Capstone C — Topic-based pub/sub

Support publish, subscribe, unsubscribe, multiple topics and slow subscribers. Choose a bounded delivery/drop policy and make send/close ownership explicit. Include concurrent lifecycle tests and metrics.

**Acceptance:** all subscribers receive accepted messages under the chosen policy; unsubscribes cannot trigger send-on-closed panics; slow subscribers do not cause unbounded queues; shutdown is idempotent.

**Relevant questions:** [Q25](25-tee-broadcast-two-outputs/README.md), [Q28](28-pub-sub/README.md), [Q29](29-detect-goroutine-leak/README.md), [Q30](30-graceful-shutdown/README.md).

The existing [14 advanced senior tasks](tasks/README.md) remain separate and intentionally unsolved; use them after the starter problems.

## Final self-check

- [ ] I can create goroutines and join them without `time.Sleep`.
- [ ] I can explain loop captures across Go language versions.
- [ ] I can predict unbuffered/buffered/nil/closed channel behavior.
- [ ] I can explain why only an owner/coordinator closes a shared output channel.
- [ ] I can find deadlock cycles and unblock with ownership/cancellation.
- [ ] I can use `select`, `context` and timeouts without assuming work was killed.
- [ ] I can identify data races and compare mutex, RWMutex and atomics.
- [ ] I can implement bounded workers, fan-out/fan-in and a cancellable pipeline.
- [ ] I can prevent send/close panics and avoid leaked goroutines.
- [ ] I can reason about rate versus concurrency, admission and backpressure.
- [ ] I can defend shutdown, error reporting, pprof evidence and happens-before.
- [ ] I can reproduce at least one failure test under `-race` and explain what that does—and does not—prove.

When revising for an interview, first answer the [30 theory prompts](THEORY-REVIEW.md) aloud, then implement the three weakest topics from a blank file.
