# 16 — Stateful designs, rate limits, and adjacent comparison

[← Foundations index](../README.md) · [← LeetCode index](../../README.md)

## Mental model
Simulation problems are easier when you list all state, establish transition rules, and distinguish accepted versus rejected operations. Sequential problems often need only the most recent accepted signature.

## Operations and reusable rules
- State is persisted across method calls using a struct and pointer receiver.
- For a logger, `lastPrinted[message]` stores the last *accepted* timestamp; rejected calls must not update it.
- For a snake game, state includes body order, occupied coordinate set, food index, score and boundaries.
- Anagram removal #2273 compares each next word's sorted key to the most recent **kept** key (not every anagram globally).
- Define a state transition table before writing logic: input, guard, mutation, return value.
- Make data structure invariants explicit: e.g. each snake body position appears in the list and set.

## Runnable example
```go
package main
import "fmt"
type Logger struct{last map[string]int}
func (l *Logger) Accept(ts int,msg string)bool{
    if prior,ok:=l.last[msg];ok && ts-prior<10{return false}
    l.last[msg]=ts // update only after accepting
    return true
}
func main(){
    l:=Logger{last:make(map[string]int)}
    fmt.Println(l.Accept(1,"foo")) // true
    fmt.Println(l.Accept(3,"foo")) // false
    fmt.Println(l.Accept(11,"foo")) // true
}
```

## Trace it on paper
`foo`: at 1, accept/store 1; at 3, reject and keep stored timestamp 1; at 11, difference=10 so accept/store 11.

## Pre-problem drill
For words `[a,b,a]` in #2273, all three remain: the two `a` entries are not consecutive anagrams after the `b`.

## Pitfalls to recall
- Do not mutate state on a rejected event unless the contract explicitly says so.
- Snake movement into the current tail is only permissible if the tail vacates on that exact move (not when eating food). Review the existing code critically.
- Specify assumptions such as nondecreasing timestamps and known direction codes.

## Problems that depend on it
- [#353 Design Snake Game](../../353%20-%20Design%20Snake%20Game/main.go)
- [#359 Logger Rate Limiter](../../359.%20Logger%20Rate%20Limiter/main.go)
- [#146 LRU Cache](../../146.%20LRU%20Cache/main.go)
- [#2273 Find Resultant Array After Removing Anagrams](../../2273.%20Find%20Resultant%20Array%20After%20Removing%20Anagrams/main.go)

## Cold recall
Write the minimum state and invariant, then rewrite the example in Go and test one edge case.
