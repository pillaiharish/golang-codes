package main

import ("reflect"; "testing")

func TestSlidingWindowMaximum(t *testing.T) {
    cases := []struct{name string; nums []int; k int; want []int}{
        {"standard", []int{1,3,-1,-3,5,3,6,7},3,[]int{3,3,5,5,6,7}},
        {"single window", []int{9,2,5},3,[]int{9}},
        {"k one", []int{4,-1,7},1,[]int{4,-1,7}},
        {"descending", []int{5,4,3,2,1},2,[]int{5,4,3,2}},
        {"ascending", []int{1,2,3,4},2,[]int{2,3,4}},
        {"equal values", []int{2,2,2},2,[]int{2,2}},
        {"empty", nil,1,[]int{}},
        {"invalid zero window", []int{1,2},0,[]int{}},
        {"invalid large window", []int{1,2},3,[]int{}},
    }
    for _,tc:=range cases {
        t.Run(tc.name,func(t *testing.T) {
            got:=maxSlidingWindow(tc.nums,tc.k)
            if !reflect.DeepEqual(got,tc.want) {t.Errorf("got %v; want %v",got,tc.want)}
        })
    }
}
