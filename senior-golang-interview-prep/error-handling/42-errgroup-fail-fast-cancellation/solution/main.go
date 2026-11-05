package main
import ("context";"errors";"fmt";"sync/atomic")
import "golang.org/x/sync/errgroup"
var ErrFailure=errors.New("task failed")
func parallel(ctx context.Context)(error,bool){g,ctx:=errgroup.WithContext(ctx);var stopped atomic.Bool
 start:=make(chan struct{})
 g.Go(func()error{<-start;return ErrFailure})
 g.Go(func()error{close(start);<-ctx.Done();stopped.Store(true);return ctx.Err()})
 err:=g.Wait();return err,stopped.Load()}
func main(){err,stopped:=parallel(context.Background());fmt.Println(err,stopped)}

