package main
import ("errors";"testing")
func TestBoth(t *testing.T){err:=both();if !errors.Is(err,ErrA)||!errors.Is(err,ErrB){t.Fatal("missing branch")}
 if errors.Unwrap(err)!=nil{t.Fatal("errors.Unwrap only handles Unwrap() error")}
 m,ok:=err.(interface{Unwrap()[]error});if !ok||len(m.Unwrap())!=2{t.Fatal("expected two children")}
}

