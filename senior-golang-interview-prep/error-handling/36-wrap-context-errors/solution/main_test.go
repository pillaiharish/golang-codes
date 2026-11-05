package main
import ("context";"errors";"testing")
func TestWrappedContext(t *testing.T){if err:=check(context.Background());err!=nil{t.Fatal(err)};ctx,cancel:=context.WithCancel(context.Background());cancel();if !errors.Is(check(ctx),context.Canceled){t.Fatal("lost cancellation")}}

