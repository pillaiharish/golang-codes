# Q14. Implement a `done` channel pattern

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../13-context-cancellation/README.md) · [Next question](../15-select-default-nonblocking/README.md)

**Level 3 — select, timeouts, and cancellation.**

## Your task

Use `done := make(chan struct{})` to broadcast a stop signal to several workers without `context.Context`.

## Scenario questions — answer before opening the solution

1. Why close rather than send one value?
2. What if two goroutines call `close(done)`?
3. Does close mean all workers finished?
4. Does a `done` channel carry errors or deadlines?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q14 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
