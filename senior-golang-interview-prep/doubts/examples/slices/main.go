package main

import (
	"fmt"
	"runtime"
)

func AppendBounded(s []int, max int, values ...int) ([]int, error) {
	if max < 0 || len(s) > max {
		return s, fmt.Errorf("invalid limit %d for length %d", max, len(s))
	}
	if len(values) > max-len(s) {
		return s, fmt.Errorf("maximum length %d exceeded", max)
	}
	return append(s, values...), nil
}

func indexingPanics() (panicked bool) {
	defer func() {
		if recover() != nil {
			panicked = true
		}
	}()
	s := make([]int, 0, 3)
	s[0] = 30 // Deliberate panic, isolated for this demonstration.
	return false
}

func check(ok bool, message string) {
	if !ok {
		panic(message)
	}
}

func main() {
	fmt.Println(runtime.Version(), runtime.GOARCH)
	panicked := indexingPanics()
	check(panicked, "zero-length indexing must panic")
	fmt.Println("zero-length index panics:", panicked)
	s := make([]int, 0, 3)
	s = s[:1]
	s[0] = 30
	fmt.Println("resliced:", s, "len", len(s), "cap", cap(s))
	s = make([]int, 0, 3)
	for i := 1; i <= 12; i++ {
		s = append(s, i)
		check(len(s) == i && cap(s) >= len(s), "invalid length/capacity")
		fmt.Println("append", i, "len", len(s), "cap", cap(s))
	}
	a := make([]int, 1, 3)
	a[0] = 10
	alias := a
	a = append(a, 20)
	a[0] = 99
	check(alias[0] == 99 && len(alias) == 1, "alias should share elements, not length")
	a = append(a, 30, 40)
	a[0] = 7
	check(alias[0] == 99, "growth should separate backing arrays")
	fmt.Println("after growth: alias", alias[0], "new", a[0])
	restricted := alias[:len(alias):len(alias)]
	larger := append(restricted, 42)
	check(len(larger) == 2 && len(restricted) == 1, "restricted capacity must allow growth")
	fmt.Println("restricted append lengths:", len(restricted), len(larger))
	bounded := make([]int, 0, 3)
	bounded, err := AppendBounded(bounded, 3, 10, 20, 30)
	check(err == nil && len(bounded) == 3, "exact limit should succeed")
	same, err := AppendBounded(bounded, 3, 40, 50)
	check(err != nil && len(same) == 3 && same[2] == 30, "overflow must reject entire batch")
	fmt.Println("bounded rejection:", err)
	_, err = AppendBounded(bounded, 3)
	check(err == nil, "empty append at limit should succeed")
	_, err = AppendBounded(nil, 0, 1)
	check(err != nil, "zero limit must reject an item")
	_, err = AppendBounded(nil, -1)
	check(err != nil, "negative limit must fail")
	_, err = AppendBounded(bounded, 2)
	check(err != nil, "already oversized input must fail")
	fmt.Println("bounded edge cases passed")
}
