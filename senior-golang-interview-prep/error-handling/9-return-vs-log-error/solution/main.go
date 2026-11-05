package main
import ("errors";"fmt";"log")
var ErrStorage = errors.New("storage offline")
func fetch()error{return ErrStorage}
func load()error{if err:=fetch();err!=nil{return fmt.Errorf("load user: %w",err)};return nil}
func main(){if err:=load();err!=nil{log.Print(err)}}

