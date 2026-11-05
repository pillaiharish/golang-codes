package main
import "testing"
func TestRecover(t *testing.T){if got:=run(true);got!="boom"{t.Fatalf("want boom got %v",got)};if run(false)!=nil{t.Fatal("unexpected recovery")}}

