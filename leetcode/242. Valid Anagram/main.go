package main

import "fmt"

// First approach
// using two maps to store the frequency of the runes and comparing them
func isAnagram(s string, t string) bool {
	if len(s) != len(t) { // check for equal length
		return false
	}
	anagram_map1 := make(map[rune]int) // map 1
	for _, v := range s {
		anagram_map1[v] += 1
	}
	anagram_map2 := make(map[rune]int) // map 2
	for _, v := range t {
		anagram_map2[v] += 1
	}

	//pass key(i.e. rune) of map1 to key of map2 to equate map1's value to map2
	for i, v := range anagram_map1 {
		if anagram_map2[i] != v {
			return false
		}
	}
	return true
}

// Second approach
// store frequency in single map while parsing string1 and decrement the same map while parsing string2
// if the map value is negative then return false else its true.
func isAnagram_two(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)

	for _, s := range s {
		count[s] += 1
	}
	for _, t := range t {
		count[t] -= 1
		if count[t] < 0 {
			return false
		}
	}
	return true
}

func main() {
	output1 := isAnagram("a", "ab")
	fmt.Println("s= \"a\" and t= \"ab\" ")
	fmt.Println("Is Anagram?")
	fmt.Println(output1)

	output2 := isAnagram_two("a", "ab")
	fmt.Println("s= \"a\" and t= \"ab\" ")
	fmt.Println("Is Anagram?")
	fmt.Println(output2)
}

// Output
// harish $ go run main.go
// s= "a" and t= "ab"
// Is Anagram?
// false
// s= "a" and t= "ab"
// Is Anagram?
// false
