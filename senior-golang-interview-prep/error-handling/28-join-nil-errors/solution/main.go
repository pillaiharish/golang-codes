package main
import ("errors";"fmt")
var ErrOne=errors.New("one")
func collected(errs ...error)error{return errors.Join(errs...)}
func main(){fmt.Println(collected(nil,nil)==nil,errors.Is(collected(nil,ErrOne),ErrOne))}

