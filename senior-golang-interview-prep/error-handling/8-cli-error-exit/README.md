# Q8. Return a nonzero CLI exit status

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../7-defer-file-close/README.md) · [Next question](../9-return-vs-log-error/README.md)

**Level:** Starter | **Concepts:** os.Exit, command boundary, stderr.

## Your task — code first

1. Accept exactly one argument and parse it as an integer.
2. Keep testable logic separate from main.
3. Return exit status 1 for invalid input and print diagnostics to stderr.

## Failure cases and acceptance tests

- [ ] Valid invocation returns status 0.
- [ ] Missing/invalid input returns status 1.

## Explain before revealing

- What is the expected behavior on success and failure?
- Which concrete error value or type should a caller be able to inspect?
- Where should this error be handled, wrapped, logged, or retried?
- What is one implementation that looks correct but loses information?

## Practice workflow

1. Implement a Go solution without opening the answer.
2. Write a failure-case test first and then the success case.
3. Compare with [the worked solution](solution/README.md); refactor your attempt.
4. Run tests and static checks, then record evidence in [revision notes](../NOTES.md).

**Mastery is not implied by the reference code.** Mark this exercise complete only after recreating and explaining it.

