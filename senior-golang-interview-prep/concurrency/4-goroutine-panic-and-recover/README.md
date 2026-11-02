# Q4. What happens if a goroutine panics?

[Concurrency index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous question](../3-goroutine-execution-order/README.md) · [Next question](../5-unbuffered-channel-send-receive/README.md)

**Level:** 1 — goroutine basics | **Focus:** panic scope, deferred functions, recovery boundary, worker failures.

## Your task

1. Start a goroutine that calls `panic("something went wrong")`; observe process exit and stack trace.
2. In a **separate run**, install a deferred `recover` inside that goroutine.
3. Make `main` wait for the worker, and explain why recovering in `main` cannot catch the worker's panic.

## Scenario questions — answer before opening the solution

1. Put `defer recover` in `main`; will it catch the worker?
2. Call `recover()` in ordinary code (not a defer); does it stop a panic?
3. Does `defer wg.Done()` prevent a crash?
4. Can a worker pool wrap each job in one recovery boundary?
5. Should every panic be ignored to keep the service alive?
6. Will `recover` catch a fatal runtime error or external process kill?

## Cold-recall prompts

- Why must Q4 use two separate programs or runs?
- What happens to subsequent statements inside a function after a panic is recovered?
- Where would you put a recovery boundary in a long-running Kafka worker, and how would you account for the failed record?
- How is a recoverable panic different from an ordinary returned `error`?

## Practice and acceptance

- [ ] Predict the blocking points, output/ordering guarantees and failure behavior before running code.
- [ ] Write your implementation from scratch, then execute it with `-race` when applicable.
- [ ] Demonstrate at least one failure or edge case and explain the synchronization mechanism.
- [ ] Record what you actually observed and what remains uncertain in [revision notes](../NOTES.md).

## Reveal the answer after your attempt

Continue to the **[Q4 worked solution, runnable code and detailed explanations](solution/README.md)**. The answer lives in `solution/` so this page remains a question-first exercise.
