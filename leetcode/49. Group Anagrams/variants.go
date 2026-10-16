package main

// groupAnagramsFreqKey is an additional approach for lowercase a-z strings.
// A fixed array is comparable and can be used directly as a map key.
// Time O(total characters); output storage O(total characters).
// The signature is NOT valid for uppercase letters or arbitrary Unicode.
func groupAnagramsFreqKey(strs []string) [][]string {
    groups:=make(map[[26]int][]string)
    for _,s:=range strs {
        var count [26]int
        for i:=0;i<len(s);i++ {
            count[s[i]-'a']++
        }
        groups[count]=append(groups[count],s)
    }
    result:=make([][]string,0,len(groups))
    for _,group:=range groups {
        result=append(result,group)
    }
    return result
}
