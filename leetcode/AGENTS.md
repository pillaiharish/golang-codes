# Rules for the `leetcode/` learning library

This file is an instruction contract for future contributors and coding assistants working inside `leetcode/`. Its goal is **understandable, revisable, repeatable, recallable** solutions, not collecting answers.

## Non-negotiable principles

1. **Inspect before editing.** Read `leetcode/README.md`, this file, `0-problem-solving-building-blocks/README.md`, and the existing problem folder. Never replace a working implementation silently.
2. **One numbered folder per LeetCode problem.** Keep existing folder names unchanged, such as `49. Group Anagrams/`; new folders use `<LeetCode number>. <Official problem title>/`. The `0-problem-solving-building-blocks/` folder is the only non-problem numbered folder.
3. **One primary, memorable solution.** Prefer the method the learner can reason about from first principles, even if a more optimized alternative exists. Put variants next to the original with descriptive names; compare the two explicitly.
4. **Preserve learning history.** If an existing solution differs from a new submission, document *both* with what changed and why. Keep existing source and demos unless a correctness change is needed, and call it out when needed.
5. **Use the same seven-part mental sequence:** recognition trigger → plain-language intuition → precise invariant → implementation recipe → hand trace → complexity/edge cases → blank-screen recall.
6. **Link concepts, not just answers.** Connect each solution to the shared building blocks, to related existing problems, and to `leetcode/README.md`.
7. **Respect problem constraints.** Be explicit about lower-case ASCII versus Unicode, sorted versus unsorted input, valid k, unique solutions, mutation of input, thread safety, and non-deterministic map iteration. Do not claim a more general solution than code supports.
8. **Test and explain failures.** Runnable Go files belong to a single `package main` per problem folder. Use table-driven `main_test.go`; `go test ./...` may not work at repository root if it is not a Go module: run `GO111MODULE=off go test` from the individual folder, or use a local module. Test edge cases and multiple variants.
9. **No fabricated progress.** Confidence, last review, and next review start blank. Do not mark a concept mastered for the learner. Capture actual outcomes in `REVISION-LOG.md`.
10. **Keep language company-neutral, clear, and falsifiable.** Never insert employer or interview-company names, or imply performance has been measured if it has not.

## Required structure

```text
leetcode/
├── AGENTS.md
├── README.md                     # master problem/pattern index, links, status
├── REVISION-LOG.md                # append-only actual revision record
├── 0-problem-solving-building-blocks/
│   └── README.md
├── <number>. <Title>/
│   ├── README.md                  # standard understanding + recall guide
│   ├── main.go                    # runnable primary and original implementations
│   ├── main_test.go               # table-driven correctness / variants
│   └── variants.go                # optional: additive approaches
```

The README must include: problem/constraints; pattern recognition clue; naïve idea versus preferred idea; invariant/proof sketch; a numbered recipe in plain English; an input-to-output trace with changing state; annotated complexity in terms of input sizes; failure traps and boundaries; alternatives and why to choose each; a 30–60-second spoken answer; and 3–5 recall prompts *before* the code. If a solution has an intentional limitation, label it near the implementation.

## Repeatable practice protocol

**Pass 1 — Understand (with notes):** State the pattern, model, and invariant; manually trace one normal case and one tricky case. Explain the code line by line.

**Pass 2 — Recall (without notes):** Hide the README and solution. Identify the pattern in 30 seconds, reconstruct the numbered recipe, and code it in 10–20 minutes. Test independently.

**Pass 3 — Transfer:** Change one assumption: duplicates, empty input, Unicode, constrained memory, invalid parameters, or streaming data. Explain whether the invariant survives.

**Pass 4 — Revisit:** Review after approximately 1 day, 3 days, 7 days, and 21 days, adjusting to performance. Failed recall restarts the cycle rather than counting as completion.

**Record reality:** After each attempt, append `YYYY-MM-DD | Problem # | Cold recall? | Correct? | Time | Confidence 1–5 | Missed invariant or syntax | Next review` to `REVISION-LOG.md`. Confidence: 1 cannot start; 2 recognizes pattern; 3 explains with notes; 4 codes from memory; 5 handles variations and defends tradeoffs.

## Adding a problem in the future

1. Inspect `leetcode/` and search for its LeetCode number (avoid duplicates).
2. Identify the shared pattern and link its building block in the new README.
3. Preserve any user-provided original solution; add improvements as labeled variants.
4. Implement and test; do not claim complexities from intuition without checking the data structure.
5. Add the problem to the index in numeric order, with its pattern, review status and relative link.
6. Do not modify unrelated legacy solutions. Raise suspected bugs as separate changes or clearly documented issues.
7. Open a PR; never merge without the learner's review.
