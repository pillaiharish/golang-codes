package main
import ("errors";"testing")
func TestFindUser(t *testing.T){u,err:=FindUser(1);if err!=nil||u.Name!="Ada"{t.Fatalf("%+v %v",u,err)};_,err=FindUser(99);if !errors.Is(err,ErrNotFound){t.Fatalf("wrong error: %v",err)}}

