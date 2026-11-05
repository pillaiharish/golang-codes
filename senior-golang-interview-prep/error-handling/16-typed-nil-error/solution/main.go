package main
import "fmt"
type MyError struct{msg string}
func(e *MyError)Error()string{return e.msg}
func broken()error{var e *MyError;return e}
func fixed()error{var e *MyError;if e==nil{return nil};return e}
func main(){fmt.Println(broken()==nil,fixed()==nil)}

