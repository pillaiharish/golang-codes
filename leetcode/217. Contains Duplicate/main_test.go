package main

import "testing"

func TestContainsDuplicate(t *testing.T) {
    cases := []struct{name string; input []int; want bool}{
        {"duplicate", []int{1,2,3,1}, true},
        {"unique", []int{1,2,3,4}, false},
        {"empty", nil, false},
        {"one", []int{42}, false},
        {"negative duplicate", []int{-4,2,-4}, true},
        {"zero duplicate", []int{0,0}, true},
    }
    for _,tc := range cases {
        t.Run(tc.name,func(t *testing.T){
            if got:=containsDuplicate(tc.input);got!=tc.want {
                t.Errorf("containsDuplicate(%v) = %v; want %v", tc.input,got,tc.want)
            }
        })
    }
}
