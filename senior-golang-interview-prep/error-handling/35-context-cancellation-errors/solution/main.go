package main
import ("context";"fmt";"time")
func DoWork(ctx context.Context)error{select{case <-time.After(time.Second):return nil;case <-ctx.Done():return ctx.Err()}}
func main(){ctx,cancel:=context.WithCancel(context.Background());cancel();fmt.Println(DoWork(ctx))}

