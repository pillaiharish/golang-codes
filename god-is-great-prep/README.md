# God Is Great — permanent Go interview-prep curriculum

This folder is a long-lived **revision system**, not a new plan for each interview.

## Start here

| Resource | What it contains |
|---|---|
| [Learning mandate](STUDY-MANDATE.md) | Definition of done, active recall, example/test standard and boundaries |
| [Layer 1 — Syntax and beginner fundamentals](layer1/README.md) | 18 beginner topics, links to full preserved explanations and checkpoint |
| [Layer 2 — Data structures and Go building blocks](layer2/README.md) | 18-topic checklist, concrete worked examples and runnable tests |
| [Layer 2 worked examples](layer2/WORKED-EXAMPLES.md) | Example, why it works, complexity, edge cases for every topic |
| [Existing Layer 1 full reference](layer1/reference.md) | Original detailed study material (preserved verbatim) |
| [Existing Layer 2 full reference](layer2/reference.md) | Original detailed study material (preserved verbatim) |

**Order:** learn syntax → implement basic collections → revisit traps in Layer 3 → reusable operations in Layer 4 → production concurrency/runtime in Layer 5 → DSA/HLD/LLD. Layers are cumulative; each review should retrieve earlier concepts from memory.

## Daily revision rule

1. Pick due concepts: Day 1, 3, 7, 14, 30 after learning (adjust when missed).
2. Close the notes; explain aloud and write a fresh example.
3. Run code/tests; change inputs to break assumptions.
4. Record exactly what failed, not just “revised”.
5. During interviews, revise weak concepts first; never reread an entire layer end-to-end.

## Out of scope

This PR covers Layer 1 and Layer 2 foundations only. Existing [senior Go material](../senior-golang-interview-prep/README.md) and [DSA building blocks](../leetcode/0-problem-solving-building-blocks/README.md) remain as more advanced references.
