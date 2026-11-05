# Q1. Implement the built-in error interface

[Error-handling index](../README.md) · [Open worked solution](solution/README.md) · [Revision notes](../NOTES.md) · [Next question](../2-divide-return-error/README.md)

**Level:** Starter | **Concepts:** error interface, method sets, pointer receivers.

## Your task — code first

1. Write the definition of the built-in error interface from memory.
2. Implement a custom error type and return it as error.
3. Explain why a single Error() string method enables decoupled callers.

## Failure cases and acceptance tests

- [ ] The custom value satisfies error.
- [ ] Its Error() string is returned as expected.

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

