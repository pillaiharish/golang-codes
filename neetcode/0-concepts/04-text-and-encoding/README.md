# Strings, bytes, runes, and representation

[All concepts](../README.md) · [150-problem index](../../README.md)

## Basic model and Go operations

A string is immutable bytes. Indexing gives a byte, while range decodes UTF-8 code points and yields byte offsets. A rune is int32, not a multi-byte UTF-8 packet. For an ASCII-only contract, byte indexing is appropriate and cheaper than materializing rune slices. Do not silently generalize ASCII letter arithmetic to Unicode.

Distinguish a substring (contiguous), a subsequence (ordered positions), and a multiset (counts with no ordering). Numeric text needs explicit parsing; string(65) encodes the code point A rather than decimal digits. Reversible encodings must preserve empty components, boundaries, and any characters that might otherwise be mistaken for separators. That is an input/format concern; choose your own codec algorithm in the problem.

## Runnable micro-example

[main.go](main.go) demonstrates only this building block. From this concept directory run `go run main.go`.

```go
s := "Aé"
fmt.Println(len(s), utf8.RuneCountInString(s))
for i, r := range s { fmt.Printf("%d %U\n", i, r) }
fmt.Println(strconv.Itoa(65))
```

The snippet is the body of the complete linked program, which includes required imports and supporting types.

## Hand trace

Expected observation: **3 2; offsets/code points 0 U+0041 and 1 U+00E9; decimal text 65**. Before running, write the value of each named variable after every statement. Check which operations copy a value and which mutate shared storage. Explain the output in terms of that state, not by memorizing it.

## Mini exercise

Compare the byte and rune counts of a combining-accent string; then state which interpretation an ASCII problem requires.

## Invariant and complexity practice

Write what must remain true before and after each operation. Count actual loop iterations, allocations, and retained references in the small example. Separate the cost of one operation from a sequence of operations. Do not assume that the linked interview problems have the same complexity as this micro-example.

## Common mistakes

Splitting UTF-8 at arbitrary bytes; conflating substring/subsequence; indexing an empty string; confusing integer conversion and parsing.

## Questions that use this building block

- [#242 Valid Anagram](../../242.%20Valid%20Anagram/README.md)
- [#49 Group Anagrams](../../49.%20Group%20Anagrams/README.md)
- [#271 Encode and Decode Strings](../../271.%20Encode%20and%20Decode%20Strings/README.md)
- [#125 Valid Palindrome](../../125.%20Valid%20Palindrome/README.md)
- [#3 Longest Substring Without Repeating Characters](../../3.%20Longest%20Substring%20Without%20Repeating%20Characters/README.md)
- [#424 Longest Repeating Character Replacement](../../424.%20Longest%20Repeating%20Character%20Replacement/README.md)
- [#567 Permutation in String](../../567.%20Permutation%20in%20String/README.md)
- [#76 Minimum Window Substring](../../76.%20Minimum%20Window%20Substring/README.md)
- [#20 Valid Parentheses](../../20.%20Valid%20Parentheses/README.md)
- [#150 Evaluate Reverse Polish Notation](../../150.%20Evaluate%20Reverse%20Polish%20Notation/README.md)
- [#297 Serialize and Deserialize Binary Tree](../../297.%20Serialize%20and%20Deserialize%20Binary%20Tree/README.md)
- [#208 Implement Trie (Prefix Tree)](../../208.%20Implement%20Trie%20%28Prefix%20Tree%29/README.md)
- [#211 Design Add and Search Words Data Structure](../../211.%20Design%20Add%20and%20Search%20Words%20Data%20Structure/README.md)
- [#212 Word Search II](../../212.%20Word%20Search%20II/README.md)
- [#131 Palindrome Partitioning](../../131.%20Palindrome%20Partitioning/README.md)
- [#17 Letter Combinations of a Phone Number](../../17.%20Letter%20Combinations%20of%20a%20Phone%20Number/README.md)
- [#127 Word Ladder](../../127.%20Word%20Ladder/README.md)
- [#332 Reconstruct Itinerary](../../332.%20Reconstruct%20Itinerary/README.md)
- [#269 Alien Dictionary](../../269.%20Alien%20Dictionary/README.md)
- [#5 Longest Palindromic Substring](../../5.%20Longest%20Palindromic%20Substring/README.md)
- [#647 Palindromic Substrings](../../647.%20Palindromic%20Substrings/README.md)
- [#91 Decode Ways](../../91.%20Decode%20Ways/README.md)
- [#139 Word Break](../../139.%20Word%20Break/README.md)
- [#1143 Longest Common Subsequence](../../1143.%20Longest%20Common%20Subsequence/README.md)
- [#97 Interleaving String](../../97.%20Interleaving%20String/README.md)
- [#115 Distinct Subsequences](../../115.%20Distinct%20Subsequences/README.md)
- [#72 Edit Distance](../../72.%20Edit%20Distance/README.md)
- [#10 Regular Expression Matching](../../10.%20Regular%20Expression%20Matching/README.md)
- [#763 Partition Labels](../../763.%20Partition%20Labels/README.md)
- [#678 Valid Parenthesis String](../../678.%20Valid%20Parenthesis%20String/README.md)
- [#43 Multiply Strings](../../43.%20Multiply%20Strings/README.md)

## Optional existing foundation lesson

[Related Go mechanics](../../../leetcode/0-problem-solving-building-blocks/04-strings-bytes-runes-and-parsing/README.md). The older library can contain worked problem solutions; stay with this lesson if you want to avoid those spoilers.

## Blank-screen recall

Recreate the micro-example, predict one changed input, and explain what information the representation preserves. Confidence and review dates remain blank until your own attempt.
