# 242 — Valid Anagram

## Before you solve: prerequisite basics

Study these in sequence and type the tiny Go example in each before reopening this solution:

1. [03 maps sets frequency and grouping](../0-problem-solving-building-blocks/03-maps-sets-frequency-and-grouping/README.md)
2. [04 strings bytes runes and parsing](../0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md)
3. [13 tries and fixed alphabet indexing](../0-problem-solving-building-blocks/13-tries-and-fixed-alphabet-indexing/README.md)

[Full mapping for all 27 LeetCode problems](../0-problem-solving-building-blocks/README.md#b-problem--exact-prerequisites-all-27-problem-folders) · [Practice log](../REVISION-LOG.md)

[← Index](../README.md) · [Building block: frequency signatures](../0-problem-solving-building-blocks/README.md#3-frequency-counts-are-a-signature-for-anagrams) · [Existing solutions](main.go) · [New [26]int variant](variants.go) · [Tests](main_test.go) · [Revision log](../REVISION-LOG.md) · [LeetCode](https://leetcode.com/problems/valid-anagram/)

## Recognition

**Clue:** "Do these strings contain exactly the same characters, with the same multiplicities, in any order?" Character order does not matter, but character counts do.

**Mental trigger:** **Same length + zero net frequency differences = anagram.**

## Intuition

A character from the first string adds one to its count; the same character from the second string subtracts one. If the two have the same multiset of characters, every counter ends at zero.

## Existing solutions — preserved

The repository's existing `main.go` has **two rune-map approaches**:

1. `isAnagram`: count runes in **two** `map[rune]int` dictionaries, then compare.
2. `isAnagram_two`: count runes from s into **one** map; decrement as runes of t arrive; if a value becomes negative, return false. Since lengths are already checked, if no count goes negative, the multisets match.

These are more general for Unicode code points than a fixed English-letter array. They do **not** perform Unicode normalization; visually identical composed/decomposed sequences need special handling outside this task.

## New approach — `[26]int` (lowercase English only)

Your new method is added as `isAnagramLowercase26` in [variants.go](variants.go). It matches the lowercase a-z LeetCode constraint and avoids hash map lookups.

## Invariant

After processing the first i bytes of **both** lowercase strings, `count[c]` equals occurrences of letter c in the processed prefix of s **minus** occurrences in the processed prefix of t.

At the end, every value must be zero.

## Rebuild it in six steps

1. If lengths differ → false.
2. Create `var count [26]int`.
3. For every byte position i, `count[s[i]-'a']++`.
4. At the same i, `count[t[i]-'a']--`.
5. Check all 26 cells; nonzero means false.
6. Otherwise return true.

### Trace: `s="aab"`, `t="aba"`

| i | s[i] | t[i] | Nonzero balances after update |
|---|---|---|---|
| 0 | a | a | none |
| 1 | a | b | a:+1, b:-1 |
| 2 | b | a | none → true |

### Counterexample: `s="aab"`, `t="abb"`

After all updates, a:+1 and b:-1. The strings have equal length but are not anagrams.

## Complexity and contract

- Fixed-array version: O(n+26)=O(n) time, O(1) extra space.
- Rune-map version: O(n) expected time, O(u) space for distinct runes (bounded by the input's character variety).
- `len(s)` counts **bytes** in Go, not Unicode code points. With the lowercase-ASCII assumption one byte = one letter. Outside that contract `s[i]-'a'` can index out of bounds; choose the rune-map version.
- Checking two strings of different byte length early is safe for exact codepoint-multiset anagrams, but byte and rune lengths are not interchangeable for arbitrary text processing.

## Why not sorting?

Sorting characters takes O(n log n), creates a sorted signature and is very useful in **Group Anagrams #49**. But checking just two lowercase strings only requires equal character frequencies; the counter is simpler and linear.

## 30-second spoken answer

> Two words are anagrams if their character frequencies match. For lowercase English letters, I use a 26-element array. After checking equal lengths, I increment for each byte from the first string and decrement for the corresponding byte of the second. If every count returns to zero, the strings are anagrams. It's O(n) time and O(1) extra space. For arbitrary Unicode, I'd use rune counters instead.

## Cold recall

1. Why doesn't equal length alone imply an anagram?
2. Why is `[26]int` sufficient for the stated constraints?
3. What does negative/positive count mean at an intermediate step?
4. How do the **two-map**, **one-map**, and **array** variants differ?
5. When does indexing `s[i]-'a'` become unsafe?

**Personal review:** Confidence `— / 5` · Last reviewed `—` · Next review `—`.
