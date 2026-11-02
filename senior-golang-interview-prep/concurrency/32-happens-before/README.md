# Q32. Explain happens-before in your code

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../31-concurrent-url-checker/README.md) · [Next question](../33-goroutine-pprof/README.md)

**Level 6 — correctness and the Go memory model.**

## Your task

For earlier exercises, explain how one goroutine knows another goroutine's writes are visible. Identify the *precise synchronization operation* creating each happens-before edge. Find an unsynchronized example and fix it.

## Scenario questions — answer before opening the solution

1. Is a scheduling delay a happens-before edge?
2. Does `WaitGroup` make multiple concurrent writes safe?
3. Does an atomic counter protect a related map?
4. Why do we care beyond the race detector?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q32 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
