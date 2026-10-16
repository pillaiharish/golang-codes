package main

import "testing"

func TestFirstMissingPositive(t *testing.T) {
    cases:=[]struct{name string; nums []int; want int}{
        {"empty",nil,1},
        {"example one",[]int{1,2,0},3},
        {"example two",[]int{3,4,-1,1},2},
        {"all bigger",[]int{7,8,9,11,12},1},
        {"already complete",[]int{1,2,3},4},
        {"duplicates",[]int{1,1},2},
        {"scrambled",[]int{2,3,1},4},
        {"negative and zero",[]int{-2,0,-1},1},
        {"n in wrong slot",[]int{3,1,2},4},
    }
    for _,tc:=range cases {
        t.Run(tc.name,func(t *testing.T) {
            got:=firstMissingPositive(append([]int(nil),tc.nums...))
            if got!=tc.want {t.Errorf("got %d, want %d",got,tc.want)}
        })
    }
}

// Compare against the obvious O(n²) oracle across many short arrays.
func TestFirstMissingPositiveAgainstBruteForce(t *testing.T) {
    values:=[]int{-1,0,1,2,3,4}
    for n:=0;n<=5;n++ {
        a:=make([]int,n)
        var walk func(int)
        walk=func(at int) {
            if at==n {
                want:=1
                for {
                    found:=false
                    for _,v:=range a {if v==want {found=true;break}}
                    if !found {break}
                    want++
                }
                got:=firstMissingPositive(append([]int(nil),a...))
                if got!=want {t.Fatalf("input %v: got %d want %d",a,got,want)}
                return
            }
            for _,v:=range values {a[at]=v;walk(at+1)}
        }
        walk(0)
    }
}
