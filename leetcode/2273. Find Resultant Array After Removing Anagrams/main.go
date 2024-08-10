package main

import (
	"fmt"
	"sort"
	"strings"
)

// My Approach 1 --- Slow, more memory, inefficient
// -----------------------------------------------------------------//
func removeAnagrams1(words []string) []string {
	result := []string{}
	mapA := make(map[string]string)
	result = append(result, words[0])
	for _, w := range words {

		temp := strings.Split(w, "")
		sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
		ws := strings.Join(temp, "")
		mapA[w] = ws
	}
	fmt.Println(mapA)
	previous := mapA[result[0]]
	for i := 1; i < len(words); i++ {
		if previous != mapA[words[i]] {
			result = append(result, words[i])
			previous = mapA[words[i]]
		}
	}
	return result
}

// My Approach 2
// -----------------------------------------------------------------//
func sortAlphabets(word string) string {
	temp := strings.Split(word, "")
	sort.Slice(temp, func(i, j int) bool { return temp[i] < temp[j] })
	return strings.Join(temp, "")
}
func removeAnagrams2(words []string) []string {
	result := []string{words[0]}

	lastSorted := sortAlphabets(words[0]) // [i-1]
	for i := 1; i < len(words); i++ {

		currentSorted := sortAlphabets(words[i]) // [i]
		// If the currentSorted word is not an anagram of the lastSorted one, add to result
		if currentSorted != lastSorted {
			result = append(result, words[i])
			lastSorted = sortAlphabets(words[i]) // currentSorted
		}

	}

	return result
}

func main() {
	// Test cases
	fmt.Println(removeAnagrams1([]string{"abba", "baba", "bbaa", "cd", "cd"})) // Output: ["abba", "cd"]
	fmt.Println(removeAnagrams1([]string{"a", "b", "c", "d", "e"}))            // Output: ["a", "b", "c", "d", "e"]

	fmt.Println(removeAnagrams2([]string{"abba", "baba", "bbaa", "cd", "cd"})) // Output: ["abba", "cd"]
	fmt.Println(removeAnagrams2([]string{"a", "b", "c", "d", "e"}))            // Output: ["a", "b", "c", "d", "e"]

}

// Output:
// map[abba:aabb baba:aabb bbaa:aabb cd:cd]
// [abba cd]
// map[a:a b:b c:c d:d e:e]
// [a b c d e]
// [abba cd]
// [a b c d e]

//-----My Hacked apporach which passed for 194 / 201 testcases passed---------------//
//
// func removeAnagrams(words []string) []string {
//     mapWords := make(map[string][]int)
//     for index, s := range words{
//         sorted := strings.Split(s,"")
//         sort.Slice(sorted, func(i,j int)bool{return sorted[i]<sorted[j]})
//         sNew := (strings.Join(sorted,""))
//         fmt.Println(sNew)
//         mapWords[sNew]= append(mapWords[sNew],index)
//     }
//     fmt.Println(mapWords)
//     count := 0
//     output:=[]int{}
//     newS :=[]string{}
//     for _, value := range mapWords{
//         count++
//        output= append(output,value[0])
//     }
//     sort.Ints(output)
//     for _, v :=range output{
//         newS=append(newS,words[v])
//     }
//     if count==len(words){
//         return words
//     }
//     return newS//[]string{"a"}
// }
// TC no 194: words = ["a","b","a"] Expected: ["a","b","a"]
