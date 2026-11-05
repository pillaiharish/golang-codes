package main
import ("fmt";"strconv")
func ParseAge(s string)(int,error){n,err:=strconv.Atoi(s);if err!=nil{return 0,fmt.Errorf("parse age: %w",err)};return n,nil}
func main(){fmt.Println(ParseAge("41"))}

