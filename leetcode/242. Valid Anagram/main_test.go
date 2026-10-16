package main

import "testing"

func TestValidAnagramVariants(t *testing.T) {
    cases:=[]struct{name,s,t string; want bool}{
        {"anagram","anagram","nagaram",true},
        {"not anagram","rat","car",false},
        {"different length","a","ab",false},
        {"empty","","",true},
        {"same","aaa","aaa",true},
        {"same length different counts","aab","abb",false},
        {"repeated","aaab","abaa",true},
    }
    variants:=[]struct{name string; fn func(string,string)bool}{
        {"two rune maps",isAnagram},
        {"one rune map",isAnagram_two},
        {"array 26",isAnagramLowercase26},
    }
    for _,tc:=range cases {
        for _,imp:=range variants {
            t.Run(tc.name+"/"+imp.name,func(t *testing.T) {
                if got:=imp.fn(tc.s,tc.t);got!=tc.want {t.Errorf("%s(%q,%q)=%v, want %v",imp.name,tc.s,tc.t,got,tc.want)}
            })
        }
    }
}

func TestExistingRuneMapUnicode(t *testing.T) {
    // Demonstrates why the lowercase-specific array variant has a separate name.
    if !isAnagram("éa","aé") || !isAnagram_two("éa","aé") {
        t.Fatal("rune map variant should work with these Unicode code points")
    }
}
