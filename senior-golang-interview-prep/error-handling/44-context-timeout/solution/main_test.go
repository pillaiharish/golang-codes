package main
import ("context";"errors";"testing";"time")
func TestTimeout(t *testing.T){if err:=slow(context.Background(),0);err!=nil{t.Fatal(err)}
 ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond);defer cancel()
 if !errors.Is(slow(ctx,time.Second),context.DeadlineExceeded){t.Fatal("missing deadline")}
}

