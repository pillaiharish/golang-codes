package main

import (
	"fmt"
	"sort"
	"strings"
)

// First approach
func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, s := range strs {
		sNew := strings.Split(s, "")
		sort.Slice(sNew, func(i, j int) bool { return sNew[i] < sNew[j] })
		sorted := strings.Join(sNew, "")
		strMap[sorted] = append(strMap[sorted], s)
	}
	fmt.Println(strMap)
	output := make([][]string, 0)
	for _, value := range strMap {
		output = append(output, value)
	}
	fmt.Println(output)
	return output //[][]string{{}}
}

// ---------------------------------------------------------
// Second approach
// Helper function to sort a string
func sortString(s string) string {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)
	return strings.Join(sorted, "")
}

func groupAnagrams2(strs []string) [][]string {
	// Map to group anagrams by their sorted string representation
	anagramGroups := make(map[string][]string)

	// Iterate over each string in the input slice
	for _, str := range strs {
		sortedStr := sortString(str)
		anagramGroups[sortedStr] = append(anagramGroups[sortedStr], str)
	}

	// Collect the groups of anagrams into a slice
	result := make([][]string, 0, len(anagramGroups))
	for _, group := range anagramGroups {
		result = append(result, group)
	}

	return result
}

func main() {
	// Test case
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(input)
	fmt.Println(result)
}

// Output:
// harish $ go run main.go
// map[abt:[bat] aet:[eat tea ate] ant:[tan nat]]
// [[eat tea ate] [tan nat] [bat]]
// [[eat tea ate] [tan nat] [bat]]
