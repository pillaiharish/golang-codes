package main
import ("errors";"strings";"testing")
func TestNoNoise(t *testing.T){err:=handler();if !errors.Is(err,ErrDenied)||strings.Count(err.Error(),"save profile")!=1{t.Fatal(err)}}

