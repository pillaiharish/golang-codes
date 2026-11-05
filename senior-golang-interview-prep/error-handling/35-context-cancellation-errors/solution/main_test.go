package main
import ("context";"errors";"testing";"time")
func TestCancellation(t *testing.T){ctx,cancel:=context.WithCancel(context.Background());cancel();if !errors.Is(DoWork(ctx),context.Canceled){t.Fatal("missing cancel")}
 d,stop:=context.WithDeadline(context.Background(),time.Now().Add(-time.Second));defer stop();if !errors.Is(DoWork(d),context.DeadlineExceeded){t.Fatal("missing deadline")}
}

