package main
import "testing"
func TestMustParseInt(t *testing.T){if MustParseInt("9")!=9{t.Fatal("bad integer")}
 defer func(){if recover()==nil{t.Fatal("expected panic")}}();MustParseInt("bad")
}

