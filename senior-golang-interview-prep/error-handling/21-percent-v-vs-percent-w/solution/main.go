package main
import ("errors";"fmt")
var ErrTarget=errors.New("target")
func wrapV(err error)error{return fmt.Errorf("op: %v",err)}
func wrapW(err error)error{return fmt.Errorf("op: %w",err)}
func main(){fmt.Println(errors.Is(wrapW(ErrTarget),ErrTarget))}

