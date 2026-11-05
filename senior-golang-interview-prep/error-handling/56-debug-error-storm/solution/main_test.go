package main
import "testing"
func TestSampling(t *testing.T){for _,tc:=range []struct{n int;want bool}{{1,true},{2,false},{99,false},{100,true},{200,true}}{if shouldLog(tc.n)!=tc.want{t.Fatalf("%d",tc.n)}}}

