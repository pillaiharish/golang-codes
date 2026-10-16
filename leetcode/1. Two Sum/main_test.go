package main

import "testing"

func TestTwoSum(t *testing.T) {
    cases := []struct {
        name string
        nums []int
        target int
        want []int
    }{
        {"simple", []int{2,7,11,15}, 9, []int{0,1}},
        {"duplicate values at distinct indices", []int{3,3}, 6, []int{0,1}},
        {"negative", []int{-3,4,3,90}, 0, []int{0,2}},
        {"zero and positive", []int{0,4,3,0}, 0, []int{0,3}},
        {"not found outside LeetCode contract", []int{1,2}, 100, nil},
    }
    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
            got := twoSum(tc.nums, tc.target)
            if len(got) != len(tc.want) { t.Fatalf("got %v, want %v",got,tc.want) }
            for i := range got { if got[i] != tc.want[i] { t.Fatalf("got %v, want %v",got,tc.want) } }
        })
    }
}
