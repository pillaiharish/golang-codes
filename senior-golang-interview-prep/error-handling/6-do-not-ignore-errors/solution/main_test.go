package main
import "testing"
func TestSafeParse(t *testing.T){if n,err:=safeParse("8");err!=nil||n!=8{t.Fatalf("%d %v",n,err)};if _,err:=safeParse("bad");err==nil{t.Fatal("invalid input accepted")}}

