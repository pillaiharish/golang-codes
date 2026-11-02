# Q13. Implement cancellation using `context.Context`

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../12-timeout-with-time-after/README.md) · [Next question](../14-done-channel-cancellation/README.md)

**Level 3 — select, timeouts, and cancellation.**

## Your task

Start a long-running worker, cancel it from `main` after about one second, and wait until it exits. Use `context.WithCancel` or `context.WithTimeout`.

## Scenario questions — answer before opening the solution

1. If the worker ignores `ctx.Done()`?
2. What if a worker is blocked sending to a channel?
3. Why call `cancel()` even with a deadline?
4. Can two child contexts share one parent deadline?
5. How should a library treat context?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q13 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
