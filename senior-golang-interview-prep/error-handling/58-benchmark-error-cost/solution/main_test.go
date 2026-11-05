package main
import ("errors";"testing")
var sink error
func TestCostContracts(t *testing.T){if !errors.Is(sentinel(),ErrMissing)||!errors.Is(contextual("a"),ErrMissing){t.Fatal("missing error")}}
func BenchmarkSentinel(b *testing.B){for i:=0;i<b.N;i++{sink=sentinel()}}
func BenchmarkContextual(b *testing.B){for i:=0;i<b.N;i++{sink=contextual("user42")}}

