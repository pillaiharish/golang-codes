package main
import ("context";"fmt";"time")
func slow(ctx context.Context,d time.Duration)error{timer:=time.NewTimer(d);defer timer.Stop();select{case <-timer.C:return nil;case <-ctx.Done():return fmt.Errorf("slow operation: %w",ctx.Err())}}
func main(){ctx,cancel:=context.WithTimeout(context.Background(),time.Millisecond);defer cancel();fmt.Println(slow(ctx,time.Second))}

