package main
import "testing"
func TestMustDivide(t *testing.T){if MustDivide(10,2)!=5{t.Fatal("wrong result")}
 defer func(){if recover()==nil{t.Fatal("expected panic")}}();MustDivide(1,0)
}

