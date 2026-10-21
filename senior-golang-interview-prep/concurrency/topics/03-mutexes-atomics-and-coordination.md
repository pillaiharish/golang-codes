# 03. Mutexes, atomics, and coordination

[Concurrency index](../README.md) · [Topic index](README.md) · [Practice rubric](../ASSESSMENT.md)

## Building blocks and senior reasoning

Choose synchronization from ownership and invariants. A mutex is a good starting point for shared mutable state. Keep network calls and arbitrary callbacks out of critical sections when possible. Define lock ordering if acquiring multiple locks. RWMutex is not automatically faster: reader coordination and writer behavior can outweigh short critical-section savings.

Atomics suit small independent state transitions and immutable pointer publication; they do not remove the need for a state-machine design. A concurrent map does not automatically make read-check-write sequences atomic. Avoid copying structs containing mutexes or other synchronization primitives after use; pointer receivers alone do not prevent callers from copying the value.

For a WaitGroup, register work before a concurrent Wait can observe a zero count; calling Add only inside the new goroutine can make Wait return too soon. Done belongs on every completion path. A WaitGroup joins work but does not cancel it or carry errors. Cond.Wait releases and reacquires its lock: check the predicate in a loop. Once models one initialization attempt, not an automatic retry policy. Pool is an allocation optimization, not a durable cache or bounded resource pool.

## Interview questions

These are representative practice prompts, not a claim that a named employer asks these exact questions. Try answering aloud before reading the answer outline.

### Q1. Mutex or channel for a shared counter map?

**Medium. Strong answer:** Start with a mutex for straightforward shared state; use an owner goroutine when command serialization helps the wider protocol.

### Q2. Why did switching to RWMutex reduce throughput?

**Hard. Strong answer:** Measure critical-section duration, contention and writer load; reader bookkeeping has a cost.

### Q3. Where should WaitGroup.Add occur?

**Hard. Strong answer:** Before work starts and before a possible Wait on zero; describe lifecycle and reuse barriers.

### Q4. Can sync.Map implement a check-then-charge transaction?

**Hard. Strong answer:** Not through unrelated operations. Define an atomic operation or protect the complete account invariant.

### Q5. Can sync.Pool cap concurrent database connections?

**Expert. Strong answer:** No. Use an explicit pool or semaphore with ownership, release and cancellation semantics.

## Implement and demonstrate

- [Task 01: Concurrent state and immutable snapshots](../tasks/01-shared-state-and-snapshots/README.md)
- [Task 03: Bounded worker pool](../tasks/03-bounded-worker-pool/README.md)
- [Task 05: Cache loading with shared work](../tasks/05-cache-and-singleflight/README.md)

## Recall checkpoint

Explain one invariant, one failure schedule, and one test that would falsify your design. Record the gap and evidence in [NOTES.md](../NOTES.md); leave mastery unclaimed until you have done the work.

## Primary references

- [https://pkg.go.dev/sync](https://pkg.go.dev/sync)
- [https://pkg.go.dev/sync/atomic](https://pkg.go.dev/sync/atomic)
