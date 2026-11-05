# Q7. Use defer and handle file cleanup

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Previous](../6-do-not-ignore-errors/README.md) · [Next question](../8-cli-error-exit/README.md)

**Level:** Starter | **Concepts:** defer, os.Open, close errors, cleanup policy.

## Your task — code first

1. Open a file and defer Close immediately after success.
2. Read its contents and propagate read errors.
3. Explain when closing a writer must be checked and why closing a read-only file often needs only best-effort cleanup.

## Failure cases and acceptance tests

- [ ] Temporary file is read correctly.
- [ ] Missing file is wrapped and remains inspectable.

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

