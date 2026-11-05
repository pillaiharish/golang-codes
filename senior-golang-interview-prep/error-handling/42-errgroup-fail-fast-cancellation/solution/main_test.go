package main
import ("context";"errors";"testing")
func TestErrgroup(t *testing.T){err,stopped:=parallel(context.Background());if !errors.Is(err,ErrFailure)||!stopped{t.Fatalf("got err=%v stopped=%v",err,stopped)}}

