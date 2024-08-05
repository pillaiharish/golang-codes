# Grouping Anagrams

## Explanation

### Helper Function (`sortString`)
This function takes a string, splits it into a slice of characters, sorts the slice, and joins it back into a string. This sorted string serves as the key for the anagram groups.

### Map (`anagramGroups`)
This map stores lists of strings, with the key being the sorted version of the string. All strings that are anagrams will have the same sorted key and will be grouped together in the map.

### Iterate and Group
For each string in the input list, we sort the string and use it as a key to append the original string to the corresponding list in the map.

### Collect Results
The values of the map (`anagramGroups`) are collected into a slice of slices, which represents the grouped anagrams.

## Complexity Analysis

### Time Complexity
Sorting each string takes \(O(m \log m)\) time, where \(m\) is the length of the string. Given \(n\) strings, the total time complexity is \(O(n * m \log m)\).

### Space Complexity
The space complexity is \(O(n * m)\), where \(n\) is the number of strings and \(m\) is the maximum length of a string, for storing the keys in the map and the grouped anagrams.