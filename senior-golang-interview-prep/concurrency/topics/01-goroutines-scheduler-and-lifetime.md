# 01. Goroutines, scheduling, and lifetime

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Concurrency means coordinating overlapping work; parallelism means executing work simultaneously. A goroutine is a runtime-managed execution context, not a dedicated OS thread. G represents a goroutine, M an OS thread, and P the runtime resources needed to execute Go code. GOMAXPROCS limits simultaneous execution of Go code, not goroutine count or all threads in the process.

Pollable network I/O normally parks a waiting goroutine while the runtime can schedule other work. A blocking system call can occupy its OS thread while other threads execute Go code. Neither mechanism makes the blocked application worker available: one reader synchronously waiting for Kafka still cannot dequeue its next record. Diagnose whether time is spent running, runnable, waiting for I/O, or waiting for synchronization before changing worker counts.

For every `go` statement, identify its owner, termination signal, blocking points and join mechanism. Cancellation without joining is not proof of termination. Goroutines can retain request bodies, buffers and other reachable data; a cheap initial stack is not a fixed lifetime memory cost. Never rely on scheduler order or sleeps for correctness.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Does one million goroutines mean one million threads?

**Medium. Strong answer:** No. Explain scheduling and retained memory, then ask what bounds admission and how all goroutines terminate.

### Q2. Why does a netpoll-aware Kafka call still bottleneck a single reader?

**Hard. Strong answer:** The OS thread can serve other goroutines, but that reader cannot advance until its synchronous call completes. Measure service time and downstream capacity.

### Q3. Will increasing GOMAXPROCS fix queue loss?

**Hard. Strong answer:** Only if CPU execution capacity is limiting. Compare CPU, runnable delay, I/O wait, throttling and queue age first.

### Q4. A handler returns but its goroutines remain: what do you inspect?

**Hard. Strong answer:** Find blocked sends, missing deadlines, retained contexts and orphan workers. Require cancellation paths and a join owned by the service or request.

### Q5. How do you distinguish a leak from a legitimate burst?

**Expert. Strong answer:** Observe counts and stack states after load stops and deadlines expire; correlate with outstanding work rather than treating one count as proof.

## Implement and demonstrate

- [Task 03: Bounded worker pool](../tasks/03-bounded-worker-pool/README.md)
- [Task 04: Cancellable fan-out and first-success](../tasks/04-cancellable-fanout/README.md)
- [Task 13: Race, deadlock, and leak diagnosis lab](../tasks/13-debugging-lab/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://go.dev/doc/diagnostics](https://go.dev/doc/diagnostics)
- [https://go.dev/src/runtime/HACKING.md](https://go.dev/src/runtime/HACKING.md)
