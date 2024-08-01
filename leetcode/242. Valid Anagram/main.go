package main

import "fmt"

// First approach
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	anagram_map1 := make(map[rune]int)
	for _, v := range s {
		anagram_map1[v] += 1
	}
	anagram_map2 := make(map[rune]int)
	for _, v := range t {
		anagram_map2[v] += 1
	}

	for i, v := range anagram_map1 {
		if anagram_map2[i] != v {
			return false
		}
	}
	return true
}

func main() {
	output := isAnagram("a", "ab")
	fmt.Println("s= \"a\" and t= \"ab\" ")
	fmt.Println("Is Anagram?")
	fmt.Println(output)
}

// Output
// harish $ go run main.go
// s= "a" and t= "ab"
// Is Anagram?
// false
