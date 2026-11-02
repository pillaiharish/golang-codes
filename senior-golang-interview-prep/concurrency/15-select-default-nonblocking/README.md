# Q15. Use `default` in `select` for nonblocking operations

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../14-done-channel-cancellation/README.md) · [Next question](../16-data-race-counter/README.md)

**Level 3 — select, timeouts, and cancellation.**

## Your task

Attempt a send only if the channel is ready to accept it; otherwise skip the value. Then implement a nonblocking receive.

## Scenario questions — answer before opening the solution

1. Does default guarantee no blocking anywhere in the branch?
2. What if you spin in a tight loop around this `select`?
3. What if the channel is closed during a send?
4. When is dropping acceptable?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q15 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
