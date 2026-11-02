# Q11. Use `select` with multiple channels

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../10-channel-deadlocks-and-lifecycle/README.md) · [Next question](../12-timeout-with-time-after/README.md)

**Level 3 — select, timeouts, and cancellation.**

## Your task

Start two senders with different delays and receive whichever channel becomes ready first. Then reason about simultaneous readiness and the sender you did not select.

## Scenario questions — answer before opening the solution

1. If both cases are ready, is the first case preferred?
2. If neither case is ready?
3. If one channel is nil?
4. If a selected channel is closed?
5. If the other sender blocks forever?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q11 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
