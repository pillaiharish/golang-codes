package main

// isAnagramLowercase26 uses one [26]int counter for the LeetCode
// lowercase-English-letters input contract. Indexing assumes 'a'..'z'.
// For general Unicode text, use the rune-map variant in main.go.
func isAnagramLowercase26(s,t string) bool {
    if len(s)!=len(t) { return false }
    var count [26]int
    for i:=0;i<len(s);i++ {
        count[s[i]-'a']++
        count[t[i]-'a']--
    }
    for _,balance:=range count {
        if balance!=0 { return false }
    }
    return true
}
