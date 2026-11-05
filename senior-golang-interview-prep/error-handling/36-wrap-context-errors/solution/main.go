package main
import ("context";"fmt")
func check(ctx context.Context)error{if err:=ctx.Err();err!=nil{return fmt.Errorf("do work: %w",err)};return nil}
func main(){ctx,cancel:=context.WithCancel(context.Background());cancel();fmt.Println(check(ctx))}

