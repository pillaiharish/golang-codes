package main
import ("context";"errors";"testing")
func TestRetry(t *testing.T){n:=0;err:=WithRetry(context.Background(),3,0,func(context.Context)error{n++;if n<3{return &RetryableError{errors.New("transient")}};return nil})
 if err!=nil||n!=3{t.Fatalf("attempts=%d err=%v",n,err)}
 permanent:=errors.New("invalid");n=0;err=WithRetry(context.Background(),4,0,func(context.Context)error{n++;return permanent});if n!=1||!errors.Is(err,permanent){t.Fatalf("permanent retried %d",n)}
 ctx,cancel:=context.WithCancel(context.Background());cancel();n=0;if !errors.Is(WithRetry(ctx,3,0,func(context.Context)error{n++;return nil}),context.Canceled)||n!=0{t.Fatal("cancel ignored")}
}

