package main

import ("reflect";"sort";"strings";"testing")

func normalizeAnagramGroups(groups [][]string) []string {
    entries:=make([]string,0,len(groups))
    for _,group:=range groups {
        copyGroup:=append([]string(nil),group...)
        sort.Strings(copyGroup)
        entries=append(entries,strings.Join(copyGroup,"\x00"))
    }
    sort.Strings(entries)
    return entries
}

func TestGroupAnagramsVariants(t *testing.T) {
    cases:=[]struct{name string; input []string; want [][]string}{
        {"classic",[]string{"eat","tea","tan","ate","nat","bat"},[][]string{{"eat","tea","ate"},{"tan","nat"},{"bat"}}},
        {"empty",[]string{},[][]string{}},
        {"empty strings",[]string{"",""},[][]string{{"",""}}},
        {"all separate",[]string{"a","b","c"},[][]string{{"a"},{"b"},{"c"}}},
        {"duplicates",[]string{"ab","ba","ab"},[][]string{{"ab","ba","ab"}}},
    }
    variants:=[]struct{name string; fn func([]string)[][]string}{
        {"original",groupAnagrams},{"original helper",groupAnagrams2},{"frequency signature",groupAnagramsFreqKey},
    }
    for _,tc:=range cases {
        for _,imp:=range variants {
            t.Run(tc.name+"/"+imp.name,func(t *testing.T) {
                got:=normalizeAnagramGroups(imp.fn(tc.input))
                want:=normalizeAnagramGroups(tc.want)
                if !reflect.DeepEqual(got,want) {
                    t.Errorf("groups=%v, want=%v",got,want)
                }
            })
        }
    }
}
