# Anagram Filter

## Explanation

### `sortString` Function
This function takes a string, splits it into characters, sorts the characters, and then joins them back into a sorted string. This sorted string representation is used to easily check for anagrams.

### Initialization
The result slice is initialized with the first word from the words array. The variable `lastSorted` keeps track of the sorted representation of the last non-anagram word added to the result.

### Iterate Through Words
For each word in the list starting from the second word, the sorted representation (referred to as `currentSorted`) is calculated. If `currentSorted` is different from `lastSorted`, the word is added to the result, and `lastSorted` is updated. This ensures that only the first word of each set of anagrams is kept.

### Return Result
Finally, the result array, which contains the words after removing anagrams, is returned.

## Complexity Analysis

### Time Complexity
The time complexity is primarily driven by the sorting operation, which is \(O(n * m \log m)\), where \(n\) is the number of words and \(m\) is the maximum length of a word.

### Space Complexity
The space complexity is \(O(n * m)\) due to the storage required for the result list and intermediate sorted representations.