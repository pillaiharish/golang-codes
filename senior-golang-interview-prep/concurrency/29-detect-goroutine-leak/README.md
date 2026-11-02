# Q29. Detect and fix a goroutine leak

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../28-pub-sub/README.md) · [Next question](../30-graceful-shutdown/README.md)

**Level 6 — production correctness.**

## Your task

Deliberately start a goroutine blocked on an unreceived channel operation. Observe goroutine counts and stack dumps; then add cancellation and prove the worker exits.

## Scenario questions — answer before opening the solution

1. Does `WaitGroup` prevent leaks?
2. Does closing a context force a blocked syscall to finish?
3. Can a leaked goroutine keep data in memory?
4. How do you test leak freedom?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q29 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
