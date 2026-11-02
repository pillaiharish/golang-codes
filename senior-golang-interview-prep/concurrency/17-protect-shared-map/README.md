# Q17. Protect a shared map

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../16-data-race-counter/README.md) · [Next question](../18-singleton-sync-once/README.md)

**Level 4 — shared state and synchronization.**

## Your task

Read and write one ordinary Go map from several goroutines. Observe the unsafe behavior with the race detector, then build versions using `sync.Mutex`, `sync.RWMutex`, and `sync.Map`.

## Scenario questions — answer before opening the solution

1. Can many goroutines read a map with no writer?
2. Does `RWMutex` always outperform `Mutex`?
3. Is `sync.Map` always ideal for read-heavy traffic?
4. What about `if _, ok := m[k]; !ok { m[k] = v }`?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q17 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
