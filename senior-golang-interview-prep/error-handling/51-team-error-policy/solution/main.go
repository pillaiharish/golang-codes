package main
import ("errors";"fmt")
var ErrConflict=errors.New("conflict")
func atBoundary(operation string,err error)error{if err==nil{return nil};return fmt.Errorf("%s: %w",operation,err)}
func main(){fmt.Println(atBoundary("create order",ErrConflict))}

