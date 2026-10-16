#!/usr/bin/env bash
# Test the nine documented Go LeetCode problems independently.
# The repository contains standalone package-main folders rather than a root go.mod.
set -euo pipefail
here="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

problems=(
  "1. Two Sum"
  "41. First Missing Positive"
  "49. Group Anagrams"
  "94. Binary Tree Inorder Traversal"
  "146. LRU Cache"
  "208. Implement Trie (Prefix Tree)"
  "217. Contains Duplicate"
  "239. Sliding Window Maximum"
  "242. Valid Anagram"
)
for problem in "${problems[@]}"; do
  printf '\n==> Testing %s\n' "$problem"
  (
    cd "$here/$problem"
    GO111MODULE=off go test -count=1 -race .
  )
done
printf '\nAll nine problem folders passed tests.\n'
