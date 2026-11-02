# Q20. Limit concurrency using a semaphore

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../19-errgroup-cancel-on-error/README.md) · [Next question](../21-worker-pool/README.md)

**Level 4 — shared state and synchronization.**

## Your task

Process many independent tasks while allowing at most **three simultaneously active tasks**. Use a buffered channel as a counting semaphore.

## Scenario questions — answer before opening the solution

1. Is a semaphore a rate limiter?
2. What if release is forgotten?
3. Is acquiring inside each goroutine equivalent?
4. What if cancellation occurs while waiting for a permit?
5. Could a fixed worker pool be simpler?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q20 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
