package main
import ("errors";"fmt")
var ErrNotFound=errors.New("not found")
func get(id int)error{return fmt.Errorf("get user %d: %w",id,ErrNotFound)}
func main(){err:=get(2);fmt.Println(err==ErrNotFound,errors.Is(err,ErrNotFound))}

