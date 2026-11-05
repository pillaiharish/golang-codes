package main
import ("context";"errors";"fmt";"math/rand";"time")
type RetryableError struct{Err error}
func(e *RetryableError)Error()string{return e.Err.Error()}
func(e *RetryableError)Unwrap()error{return e.Err}
func WithRetry(ctx context.Context,attempts int,backoff time.Duration,fn func(context.Context)error)error{
 if attempts<1{return errors.New("attempts must be positive")};if backoff<0{return errors.New("backoff must not be negative")}
 var last error
 for i:=0;i<attempts;i++{
  if err:=ctx.Err();err!=nil{return err}
  last=fn(ctx);if last==nil{return nil};var re *RetryableError
  if !errors.As(last,&re)||i==attempts-1{return last}
  d:=backoff;for step:=0;step<i&&d<time.Second;step++{d*=2};if d>time.Second{d=time.Second}
  if d>0{d+=time.Duration(rand.Int63n(int64(d)+1))}
  timer:=time.NewTimer(d)
  select{case <-ctx.Done():timer.Stop();return ctx.Err();case <-timer.C:}
 }
 return last
}
func main(){count:=0;fmt.Println(WithRetry(context.Background(),3,0,func(context.Context)error{count++;if count<3{return &RetryableError{errors.New("try again")}};return nil}))}

