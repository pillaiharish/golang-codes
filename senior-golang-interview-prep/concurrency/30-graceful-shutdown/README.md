# Q30. Implement graceful shutdown

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../29-detect-goroutine-leak/README.md) · [Next question](../31-concurrent-url-checker/README.md)

**Level 6 — production correctness.**

## Your task

Run a service with workers, listen for SIGINT, stop accepting new work, let admitted jobs finish within a bounded grace period, and then release resources. Use `signal.NotifyContext`.

## Scenario questions — answer before opening the solution

1. Why stop admission before closing a jobs channel?
2. Is canceling workers immediately always graceful?
3. What if Kafka has buffered records?
4. What if shutdown is called twice?
5. Does the deadline guarantee the process has no goroutines?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q30 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
